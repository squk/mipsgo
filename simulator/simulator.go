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

func (s *Simulator) Run() {
	s.Lexer.Lex()
	s.Lexer.PrintTokens()
	s.Parser.Parse(s.Lexer.Tokens)
	s.Parser.PrintInstructions()
}

func (s *Simulator) GetSource() {
	b, err := ioutil.ReadFile(s.Filename)

	if err != nil {
		fmt.Print(err)
	}

	s.Source = b
}
