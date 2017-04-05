package simulator

func (vm *VirtualMachine) ADD(instr Instruction) {
	var val int32 = 0
	if instr.RT == -1 {
		val = instr.Immediate
	} else {
		val = vm.Registers[instr.RT]
	}
	vm.Registers[instr.RD] = vm.Registers[instr.RS] + val
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
	var val int32 = 0
	if instr.RT == -1 {
		val = instr.Immediate
	} else {
		val = vm.Registers[instr.RT]
	}

	var result int32 = 0
	if vm.Registers[instr.RS] < val {
		result = 1
	}

	vm.Registers[instr.RD] = result
}
