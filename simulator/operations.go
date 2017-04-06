package simulator

import (
	"errors"
	"fmt"
	"strconv"
)

// instruction formats for MIPS
const (
	R = iota
	I
	J
)

func ValidateInstruction(instr Instruction, format int) error {
	message := ("Line " + strconv.Itoa(instr.LineNumber) + ": ")
	if format == R {
		if instr.Immediate != 0 || instr.RD == -1 || instr.RS == -1 || instr.RT == -1 {
			return errors.New(message + "Immediate passed instead of RT or RT is missing")
		}
	} else if format == I {
		/* according to the MIPS specification, RS and RT are the names of the
		 * registers passed to I-format instructions. However it's not worth
		 * parsing different instruction formats differently just to change the
		 * label for the argument we're passing */

		if instr.RT != -1 {

			return errors.New("Line " + string(instr.LineNumber) + ": Register passed as argument instead of immediate")
		}
	} else if format == J {
		if instr.Label == "" || instr.Immediate != 0 || instr.RD != -1 || instr.RS != -1 || instr.RT != -1 {
			return errors.New("Line " + string(instr.LineNumber) + ": Jump instruction missing label or has extra parameter")
		}
	}

	return nil
}

func (vm *VirtualMachine) ADD(instr Instruction) error {
	err := ValidateInstruction(instr, R)
	if err != nil {
		return err
	}

	vm.Registers[instr.RD] = vm.Registers[instr.RS] + vm.Registers[instr.RT]
	return nil
}

func (vm *VirtualMachine) ADDI(instr Instruction) error {
	err := ValidateInstruction(instr, I)
	if err != nil {
		return err
	}

	vm.Registers[instr.RD] = vm.Registers[instr.RS] + instr.Immediate

	return nil
}

func (vm *VirtualMachine) SUB(instr Instruction) error {
	err := ValidateInstruction(instr, R)
	if err != nil {
		return err
	}

	val := vm.Registers[instr.RT]

	vm.Registers[instr.RD] = vm.Registers[instr.RS] - val
	return nil
}

func (vm *VirtualMachine) SLL(instr Instruction) error {
	shamt := instr.Immediate
	var result int32 = vm.Registers[instr.RS] << uint32(shamt)

	vm.Registers[instr.RD] = result
	return nil
}

func (vm *VirtualMachine) SLT(instr Instruction) error {
	err := ValidateInstruction(instr, R)
	if err != nil {
		return err
	}

	var result int32 = 0
	if vm.Registers[instr.RS] < vm.Registers[instr.RT] {
		result = 1
	}

	vm.Registers[instr.RD] = result
	return nil
}

func (vm *VirtualMachine) SLTI(instr Instruction) error {
	err := ValidateInstruction(instr, I)
	if err != nil {
		return err
	}

	var result int32 = 0
	if vm.Registers[instr.RS] < instr.Immediate {
		result = 1
	}

	vm.Registers[instr.RD] = result
	return nil
}

func (vm *VirtualMachine) JUMP(instr Instruction, validate bool) error {
	if validate == true {
		err := ValidateInstruction(instr, J)
		if err != nil {
			return err
		}
	}

	if instr.Label != "" {
		for i, searchItem := range *(vm.Instructions) {
			if searchItem.OpCode == 0 && searchItem.Label == instr.Label {
				vm.PC = int32(i - 1)
			}
		}
	}

	return nil
}

func (vm *VirtualMachine) BEQ(instr Instruction) error {
	err := ValidateInstruction(instr, I)
	if err != nil {
		return err
	}

	if instr.Label != "" {
		if vm.Registers[instr.RD] == vm.Registers[instr.RS] {
			vm.JUMP(instr, false)
		}
	}

	return nil
}

func (vm *VirtualMachine) BNE(instr Instruction) error {
	err := ValidateInstruction(instr, I)
	if err != nil {
		return err
	}

	if instr.Label != "" {
		if vm.Registers[instr.RD] != vm.Registers[instr.RS] {
			vm.JUMP(instr, false)
		}
	}

	return nil
}

func (vm *VirtualMachine) SW(instr Instruction) error {
	err := ValidateInstruction(instr, I)
	if err != nil {
		return err
	}

	// SW operation in MIPS uses RT for the register to fetch its data from and
	// RS for the address. We use RD and RS respectively for ease of parsing
	address := vm.Registers[instr.RS]
	value := vm.Registers[instr.RD]
	offset := instr.Immediate

	vm.Memory.SetWord(address+offset, WORD(value))

	return nil
}

func (vm *VirtualMachine) LW(instr Instruction) error {
	err := ValidateInstruction(instr, I)
	if err != nil {
		return err
	}

	// LW operation in MIPS uses RT for the register to load data into and RS
	// for the address. We use RD and RS respectively for ease of parsing
	address := vm.Registers[instr.RS]
	offset := instr.Immediate
	vm.Registers[instr.RD] = int32(vm.Memory.GetWord(address + offset))

	return nil
}

// MIPSGO SPECIFIC PSUEDO INSTRUCTIONS
func (vm *VirtualMachine) PBIN(instr Instruction) error {
	value := vm.Registers[instr.RD]
	//fmt.Println(strconv.FormatInt(int64(value), 2))
	//fmt.Println(strconv.FormatUint(uint64(value), 2))
	fmt.Printf("%b\n", uint32(value))
	return nil
}

func (vm *VirtualMachine) PHEX(instr Instruction) error {
	value := vm.Registers[instr.RD]
	fmt.Printf("0x%x\n", uint32(value))
	return nil
}

func (vm *VirtualMachine) PDEC(instr Instruction) error {
	value := vm.Registers[instr.RD]
	fmt.Println(value)
	return nil
}
