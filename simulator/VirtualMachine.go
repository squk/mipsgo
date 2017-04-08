package simulator

import "errors"

type VirtualMachine struct {
	Registers    []int32
	Instructions *([]Instruction)
	Memory       VMem
	PC           int32
	Outputs      []string
}

var RegMap = map[string]int{
	"zero": 0,
	"at":   1,
	"v0":   2,
	"v1":   3,
	"a0":   4,
	"a1":   5,
	"a2":   6,
	"a3":   7,
	"t0":   8,
	"t1":   9,
	"t2":   10,
	"t3":   11,
	"t4":   12,
	"t5":   13,
	"t6":   14,
	"t7":   15,
	"s0":   16,
	"s1":   17,
	"s2":   18,
	"s3":   19,
	"s4":   20,
	"s5":   21,
	"s6":   22,
	"s7":   23,
	"t8":   24,
	"t9":   25,
	"k0":   26,
	"k1":   27,
	"gp":   28,
	"sp":   29,
	"fp":   30,
	"ra":   31,
}

func InitVM() VirtualMachine {
	var vm VirtualMachine
	vm.Registers = make([]int32, 32)
	vm.Outputs = make([]string, 0)
	vm.Memory = InitMemory()
	vm.PC = 0
	return vm
}

func (vm *VirtualMachine) GetReg(s string) *int32 {
	var reg *int32
	placeholder := int32(1)
	reg = &placeholder

	if val, ok := RegMap[s]; ok {
		if val < len(vm.Registers) {
			reg = &vm.Registers[val]
		}
	}

	return reg
}

func GetRegNumber(s string) int {
	if val, ok := RegMap[s]; ok {
		return val
	} else {
		return 0
	}
}

func GetRegName(n int) string {
	name := ""

	for k, v := range RegMap {
		if v == n {
			name = k
		}
	}

	return name
}

func (vm *VirtualMachine) GetMappedRegisters() map[string]int32 {
	newMap := make(map[string]int32, 32)

	for k, i := range RegMap {
		newMap[k] = vm.Registers[i]
	}

	return newMap
}

func (vm *VirtualMachine) Print(str string) {
	vm.Outputs = append(vm.Outputs, str)
}

func (vm *VirtualMachine) RunInstruction() error {
	var err error
	if vm.Instructions == nil {
		return errors.New("No instructions provided")
	}
	if int(vm.PC) >= len(*vm.Instructions) {
		return nil
	}
	instr := (*vm.Instructions)[vm.PC]

	switch operations[instr.OpCode] {
	case "noop":
		break
	case "add":
		err = vm.ADD(instr)
	case "addi":
		err = vm.ADDI(instr)
	case "sub":
		err = vm.SUB(instr)
	case "sll":
		err = vm.SLL(instr)
	case "slt":
		err = vm.SLT(instr)
	case "slti":
		err = vm.SLTI(instr)
	case "j":
		err = vm.JUMP(instr, true)
	case "beq":
		err = vm.BEQ(instr)
	case "bne":
		err = vm.BNE(instr)
	case "sw":
		err = vm.SW(instr)
	case "lw":
		err = vm.LW(instr)

	// MIPSGO pseudo instructions
	case "pbin":
		err = vm.PBIN(instr)
	case "phex":
		err = vm.PHEX(instr)
	case "pdec":
		err = vm.PDEC(instr)
	default:
		break
	}

	vm.PC++

	return err
}
