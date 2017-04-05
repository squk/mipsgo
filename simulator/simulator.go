package simulator

import (
	"fmt"
	"io/ioutil"
)

type Simulator struct {
	Filename string
	Source   []byte
	Lexer    Lexer
	Parser   Parser
	VM       VirtualMachine
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

func ReadSource(filename string) Simulator {
	var s Simulator
	s.Filename = filename
	s.GetSource()
	s.Lexer = NewLexer()
	s.Lexer.Raw = s.Source

	return s
}

func (s *Simulator) PreProcess() {
	s.Lexer.Lex()
	s.Parser.Parse(s.Lexer.Tokens)
}

func (s *Simulator) Run() {
	s.PreProcess()
	s.RunCode()
}

func (s *Simulator) RunCode() {
	instructions := &(s.Parser.Instructions)
	s.VM.Instructions = instructions

	pc := s.VM.PC
	for pc != int32(len(*instructions)) {
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
