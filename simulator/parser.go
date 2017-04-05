package simulator

import "fmt"

type Parser struct {
	Tokens       []Token
	Instructions []Instruction
}

func NewParser() Parser {
	return Parser{
		Tokens: make([]Token, 0),
	}
}

func (p *Parser) Parse(tokens []Token) {
	p.Tokens = tokens

	for i := 0; i < len(p.Tokens); {
		if p.Tokens[i].Category == KEYWORD {
			if p.Tokens[i].ID == "break" {
				// TODO: Implement a breakpoint pseudo-instruction
			} else {
				i = p.ParseOperation(i)
			}
		}
	}
}

func (p *Parser) ParseOperation(index int) int {
	var instr Instruction
	instr = NewInstruction()
	newIndex := index + 1

	instr.OpCode = GetOpCode(p.Tokens[index].ID)

	// no matching opcode found	OR noop
	if instr.OpCode == 0 {
		return newIndex
	}

	var i int

	for i = newIndex; i < len(p.Tokens); i++ {
		tk := p.Tokens[i]

		if tk.Category == SYMBOL {
			if tk.ID == "$" {
				i++
				tk = p.Tokens[i]

				if tk.Category == TEXT || (tk.Category == NUMBER && tk.ID == "0") {
					instr.AddArgument(tk.ID)
				}

			} else if tk.ID == "," {
				continue
			}
		} else if tk.Category == NUMBER {
			instr.Immediate = int32(tk.Value)
		} else if tk.Category == TEXT {
			break
		}

		if tk.HasNL {
			i++
			break
		}
	}

	p.Instructions = append(p.Instructions, instr)
	return i
}

func (p *Parser) PrintInstructions() {
	fmt.Println("Instructions: ")

	fmt.Println("\tOP  RD  RS  RT  IMM")

	for _, instr := range p.Instructions {
		fmt.Println("\t", instr.OpCode, " ", instr.RD, " ", instr.RS, " ", instr.RT, " ", instr.Immediate)
	}
}
