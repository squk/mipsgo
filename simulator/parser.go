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
	fmt.Println("parsing...")
	p.Tokens = tokens

	for i := 0; i < len(p.Tokens); {
		tk := p.Tokens[i]
		if tk.Category == KEYWORD {
			if tk.ID == "break" {
				// TODO: Implement a breakpoint pseudo-instruction
			} else {
				i = p.ParseOperation(i)
			}
		} else if tk.Category == TEXT {
			i = p.ParseLabel(i)
		}
	}
}

// because of how the stack pointer in MIPS works, we want to store our labels
// as noop instructions
func (p *Parser) ParseLabel(index int) int {
	newIndex := index + 1

	if newIndex < len(p.Tokens) {
		if p.Tokens[newIndex].ID == ":" {
			labelInstruction := NewInstruction()
			labelInstruction.Label = p.Tokens[index].ID
			labelInstruction.LineNumber = p.Tokens[index].LineNumber
			p.Instructions = append(p.Instructions, labelInstruction)
			newIndex++
		}
	}

	return newIndex
}

func (p *Parser) ParseOperation(index int) int {
	var instr Instruction
	instr = NewInstruction()
	newIndex := index + 1

	if index < len(p.Tokens) {
		instr.OpCode = GetOpCode(p.Tokens[index].ID)
		instr.LineNumber = p.Tokens[index].LineNumber
	}

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
			instr.Label = tk.ID
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

	fmt.Println("\tOP\tRD\tRS\tRT\tIMM\tLBL")

	for _, instr := range p.Instructions {
		fmt.Println("\t", instr.OpCode, "\t", instr.RD, "\t", instr.RS, "\t", instr.RT, "\t", instr.Immediate, "\t", instr.Label)
	}
}
