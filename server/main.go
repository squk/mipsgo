package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"

	. "github.com/ctnieves/mipsgo/simulator"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

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
	response      Response
}

// Commands used in Requests
const (
	RUN       = "run"
	STEP      = "step"
	WRITE_MEM = "write_memory"
	CLEAR_MEM = "clear_memory"
)

type Request struct {
	Sender  string `json:"sender,omitempty"`
	Source  string `json:"source,omitempty"`
	Command string `json:"command,omitempty"`
	Memory  string `json:"memory"`
}

type Response struct {
	RegisterContents map[string]int32 `json:"registers"`
	Output           string           `json:"output"`
	Memory           string           `json:"memory"`
	Data             struct {
		CurrentLine int `json:"current_line"`
	} `json:"data"`
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

		if req.Command == RUN || req.Command == STEP {
			if c.currentSource != req.Source {
				c.currentSource = req.Source
				c.simulator.SetSource(req.Source)
				c.simulator.Init()
			}

			go c.remoteRun(req, req.Command)
		} else if req.Command == WRITE_MEM {
			hexString := req.Memory
			c.simulator.VM.Memory.WriteMemory(hexString)
		}
	}
}

func (c *Client) remoteRun(req Request, cmd string) {
	defer c.simulator.ClearOutputs()

	if !c.simulator.Running {
		c.simulator.Init()
		c.simulator.SetSource(req.Source)
	}
	var err error = nil
	if cmd == RUN {
		err = c.simulator.Run()
		if !c.simulator.Paused {
			c.response.Output += "Run complete...\n"
		}
	} else if cmd == STEP {
		c.simulator.Step()
	}

	if err != nil {
		c.response.Output += err.Error() + "\n"
	}

	c.response.Memory = c.simulator.VM.Memory.ToText()
	c.response.RegisterContents = c.simulator.VM.GetMappedRegisters()

	for _, out := range c.simulator.VM.Outputs {
		c.response.Output += out
	}
	c.response.Output += "\n"
	c.response.Data.CurrentLine = c.simulator.GetCurrentLine()

	resp, err := json.Marshal(c.response)
	if err != nil {
		fmt.Println("Error marshalling console output for browser")
	} else {
		c.socket.WriteMessage(websocket.TextMessage, resp)
	}

	// clear response
	c.response = Response{}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case _, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
			}
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

var templates = template.Must(template.New("temps").Funcs(template.FuncMap{
	"minus": func(a, b int) int {
		return a - b
	},
	"plus": func(a, b int) int {
		return a + b
	},
	"rand": func(n int) int {
		return rand.Intn(n)
	},
	"loop": func(n int) []int {
		var arr = make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i
		}
		return arr
	},
}).ParseGlob("public/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func handleFileServers(directories []string) {
	for _, dir := range directories {
		d := http.FileServer(http.Dir("./public/" + dir))
		http.Handle("/"+dir+"/", http.StripPrefix("/"+dir+"/", d))
	}
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+filename)
	})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	go manager.start()
	handleFileServers([]string{"css", "fonts", "js", "js/ace"})

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index.html", indexHandler)
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":"+port, nil)
}
