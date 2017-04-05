package simulator

func (vm *VirtualMachine) ADD(instr Instruction) {
	vm.Registers[instr.RD] = vm.Registers[instr.RS] + vm.Registers[instr.RT]
}

func (vm *VirtualMachine) ADDI(instr Instruction) {
	vm.Registers[instr.RD] = vm.Registers[instr.RS] + instr.Immediate
}

func (vm *VirtualMachine) SUB(instr Instruction) {
	var val int32 = 0
	if instr.RT == -1 {
		val = instr.Immediate
	} else {
		val = vm.Registers[instr.RT]
	}
	vm.Registers[instr.RD] = vm.Registers[instr.RS] - val
}

func (vm *VirtualMachine) SLT(instr Instruction) {
	var result int32 = 0
	if vm.Registers[instr.RS] < vm.Registers[instr.RT] {
		result = 1
	}

	vm.Registers[instr.RD] = result
}

func (vm *VirtualMachine) SLTI(instr Instruction) {
	var result int32 = 0
	if vm.Registers[instr.RS] < instr.Immediate {
		result = 1
	}

	vm.Registers[instr.RD] = result
}
