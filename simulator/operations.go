package simulator

func (vm *VirtualMachine) add(dest, r1, r2 int) {
	vm.registers[dest] = vm.registers[r1] + vm.registers[r2]
}

func (vm *VirtualMachine) sub(dest, r1, r2 int) {
	vm.registers[dest] = vm.registers[r1] - vm.registers[r2]
}

func (vm *VirtualMachine) slt(dest, r1, r2 int) {
	var val int32 = 0
	if vm.registers[r1] < vm.registers[r2] {
		val = 1
	}
	vm.registers[dest] = val
}
