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
	sp := s.VM.Registers[GetRegNumber("sp")]

	for sp != int32(len(*instructions)) {
		s.RunInstruction((*instructions)[sp])
		s.VM.IncSP()
		sp = s.VM.Registers[GetRegNumber("sp")]

	}
}

func (s *Simulator) GetSource() {
	b, err := ioutil.ReadFile(s.Filename)

	if err != nil {
		fmt.Print(err)
	}

	s.Source = b
}

func (s *Simulator) RunInstruction(instr Instruction) {
	var err error

	switch operations[instr.OpCode] {
	case "noop":
		break
	case "add":
		err = s.VM.ADD(instr)
	case "addi":
		err = s.VM.ADDI(instr)
	case "sub":
		err = s.VM.SUB(instr)
	case "slt":
		err = s.VM.SLT(instr)
	case "slti":
		err = s.VM.SLTI(instr)
	default:
		fmt.Println("no op found")
	}

	// TODO: handle errors more gracefully than a printout
	if err != nil {
		fmt.Println(err)
	}
}
