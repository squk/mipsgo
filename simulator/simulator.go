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
	s.Parser.Parse(s.Lexer.Tokens)
	s.Lexer.PrintTokens()
	s.Parser.PrintInstructions()
}

func (s *Simulator) Run() {
	s.Running = true
	start := time.Now()

	s.PreProcess()
	s.RunCode()

	elapsed := time.Since(start)
	fmt.Printf("Parse and Run took: %s \n", elapsed)
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

func (s *Simulator) RunCode() {
	instructions := &(s.Parser.Instructions)
	s.VM.Instructions = instructions

	pc := s.VM.PC
	for pc != int32(len(*instructions)) && !s.Paused {
		s.VM.RunInstruction()
		pc = s.VM.PC
	}
}

func (s *Simulator) GetSource() {
	b, err := ioutil.ReadFile(s.Filename)

	if err != nil {
		fmt.Print(err)
	}

	s.Source = b
}
