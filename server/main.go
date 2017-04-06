package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	. "github.com/ctnieves/mipsgo/simulator"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	p := 1
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	id            string
	socket        *websocket.Conn
	send          chan []byte
	simulator     Simulator
	currentSource string
}

type Request struct {
	Sender  string `json:"sender,omitempty"`
	Source  string `json:"source,omitempty"`
	Command string `json:"command,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			conn.simulator = EmptySimulator()
			manager.clients[conn] = true
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
			}
		// sends message to all clients
		case request := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- request:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()

		// client most likely disconnected
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		req := Request{Sender: c.id}
		err = json.Unmarshal(message, &req)

		if c.currentSource != req.Source {
			c.currentSource = req.Source
			c.simulator.SetSource(req.Source)
		}

		if req.Command == "run" {
			c.simulator.Init()
			c.simulator.SetSource(req.Source)
			c.simulator.Run()
		}
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func wsPage(res http.ResponseWriter, r *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, r, nil)
	if error != nil {
		http.NotFound(res, r)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}

	manager.register <- client

	go client.read()
	go client.write()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	go manager.start()
	fs := http.FileServer(http.Dir("ace"))
	http.Handle("/ace/", http.StripPrefix("/ace/", fs))
	http.HandleFunc("/", handler)
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":"+port, nil)
}
