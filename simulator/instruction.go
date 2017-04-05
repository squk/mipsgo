package simulator

// operations index
var operations = [...]string{
	"noop", "add", "addi", "addiu", "addu", "and", "andi", "beq",
	"bgez", "bgezal", "bgtz", "blez", "bltz", "bltzal",
	"bne", "div", "divu", "j", "jal", "jr", "lb", "lui",
	"lw", "mfhi", "mflo", "mult", "multu", "or",
	"ori", "sb", "sll", "sllv", "slt", "slti", "sltiu",
	"sltu", "sra", "srl", "srlv", "sub", "subu", "sw",
	"syscall", "xor", "xori",
}

func GetOpCode(str string) int {
	opcode := 0
	for i, code := range operations {
		if code == str {
			opcode = i
		}
	}

	return opcode
}

// all keywords are operations FOR NOW: .data, .text, and
// .globl will have to be KWs as well
type Instruction struct {
	OpCode    int
	RD        int
	RS        int
	RT        int
	Shift     int
	Funct     int
	Immediate int
	Address   int
}

func NewInstruction() Instruction {
	return Instruction{
		OpCode: 0,
		RD:     -1,
		RS:     -1,
		RT:     -1,
	}
}

func (instr *Instruction) AddArgument(arg string) {
	if instr.RD == -1 {
		instr.RD = GetRegNumber(arg)
	} else if instr.RS == -1 {
		instr.RS = GetRegNumber(arg)
	} else if instr.RT == -1 {
		instr.RT = GetRegNumber(arg)
	}
}
