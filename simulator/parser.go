package simulator

import (
	"fmt"
	"strconv"
)

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
			i = p.ParseOperation(i)
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
		} else {
			// for some reason an identifier exists that isn't a label or a KW
			newIndex = p.SkipPastNL(newIndex)
		}
	}

	return newIndex
}

func (p *Parser) SkipPastNL(index int) int {
	newIndex := index

	for i := index; i < len(p.Tokens); i++ {
		if p.Tokens[i].HasNL {
			newIndex = i
		}
	}

	return newIndex + 1
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
	if instr.OpCode <= 0 {
		return newIndex
	}
	if operations[instr.OpCode] == "break" {
		p.Instructions = append(p.Instructions, instr)
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
		} else {
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

func (p *Parser) GetInstructions() string {
	str := "Instructions: \n"
	str += "\tOP\tRD\tRS\tRT\tIMM\tLBL\n"

	for _, instr := range p.Instructions {
		str += ("\t" + strconv.Itoa(instr.OpCode) + "\t" +
			strconv.Itoa(instr.RD) + "\t" + strconv.Itoa(instr.RS) + "\t" +
			strconv.Itoa(instr.RT) + "\t" + strconv.FormatInt(int64(instr.Immediate), 10) + "\t" +
			instr.Label + "\n")
	}
	return str
}
