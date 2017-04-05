package simulator

import "errors"

// instruction formats for MIPS
const (
	R = iota
	I
	J
)

func ValidateInstruction(instr Instruction, format int) error {
	if format == R {
		if instr.Immediate != 0 || instr.RD == -1 || instr.RS == -1 || instr.RT == -1 {
			return errors.New("Immediate passed instead of RT or RT is missing")
		}
	} else if format == I {
		/* according to the MIPS specification, RS and RT are the names of the
		 * registers passed to I-format instructions. However it's not worth
		 * parsing different instruction formats differently just to change the
		 * label for the argument we're passing */

		if instr.RT != -1 {
			return errors.New("Register passed as argument instead of immediate")
		}
	} else if format == J {
		// TODO: validate Jump instruction format
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

	var val int32 = 0
	if instr.RT == -1 {
		val = instr.Immediate
	} else {
		val = vm.Registers[instr.RT]
	}

	vm.Registers[instr.RD] = vm.Registers[instr.RS] - val
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
