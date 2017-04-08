package simulator

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Simulator struct {
	Filename string
	Source   []byte
	Lexer    Lexer
	Parser   Parser
	VM       VirtualMachine
	Paused   bool
	Running  bool
}

func EmptySimulator() Simulator {
	var s Simulator
	s.Filename = ""
	s.Source = []byte("")
	s.Lexer = NewLexer()
	s.Lexer.Raw = s.Source
	s.Parser = NewParser()
	s.VM = InitVM()

	return s
}

func NewSimulator(src string) Simulator {
	var s Simulator
	s.Filename = ""
	s.Source = []byte(src)
	s.Lexer = NewLexer()
	s.Lexer.Raw = s.Source
	s.Parser = NewParser()
	s.VM = InitVM()

	return s
}

func (s *Simulator) Init() {
	s.Parser = NewParser()
	s.VM = InitVM()
}

func (s *Simulator) SetSource(src string) {
	s.Source = []byte(src)
	s.Lexer = NewLexer()
	s.Lexer.Raw = s.Source
}

func (s *Simulator) PreProcess() {
	s.Lexer.Lex()
	fmt.Println(s.Lexer.GetTokens())
	s.Parser.Parse(s.Lexer.Tokens)
}

func (s *Simulator) Run() error {
	s.Running = true
	s.Paused = false
	start := time.Now()

	s.PreProcess()
	err := s.RunCode()

	if !s.Paused {
		elapsed := time.Since(start)
		fmt.Printf("Parse and Run took: %s \n", elapsed)
	}
	return err
}

func (s *Simulator) GetTokensAndInstructions() string {
	return s.Lexer.GetTokens() + s.Parser.GetInstructions()
}

func (s *Simulator) Step() {
	s.Paused = true
	if !s.Running {
		s.PreProcess()
	}

	if s.VM.PC <= int32(len(*s.VM.Instructions)) {
		s.VM.RunInstruction()
	}
}

func (s *Simulator) RunCode() error {
	fmt.Println("Running")
	instructions := &(s.Parser.Instructions)
	s.VM.Instructions = instructions

	pc := s.VM.PC
	for pc != int32(len(*instructions)) && !s.Paused {
		if operations[(*instructions)[s.VM.PC].OpCode] == "break" {
			s.Paused = true
			fmt.Println("PAUSED")
			s.VM.Outputs = append(s.VM.Outputs, "Execution paused")
			s.VM.PC++
			return nil
		} else {
			err := s.VM.RunInstruction()
			pc = s.VM.PC

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Simulator) GetSource() {
	b, err := ioutil.ReadFile(s.Filename)

	if err != nil {
		fmt.Print(err)
	}

	s.Source = b
}
