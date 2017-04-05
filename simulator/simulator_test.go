package simulator_test

import (
	. "github.com/ctnieves/mipsgo/simulator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simulator", func() {
	var sim Simulator

	It("performs ADDI op", func() {
		sim = NewSimulator(`addi $t0, $0, 5
							addi $t1, $t0, 10`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(5)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(15)))
	})

	It("performs ADD op", func() {
		sim = NewSimulator(`addi $t0, $0, 15
							addi $t1, $0, 14
							add $t2, $t0, $t1`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(15)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(14)))
		Expect(*(sim.VM.GetReg("t2"))).To(Equal(int32(29)))
	})

	It("performs SUB op", func() {
		sim = NewSimulator(`addi $t0, $0, 15
							addi $t1, $0, 14
							sub $t2, $t0, $t1`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(15)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(14)))
		Expect(*(sim.VM.GetReg("t2"))).To(Equal(int32(1)))
	})

	It("performs SLT OP", func() {
		sim = NewSimulator(`addi $t0, $0, 5
							addi $t1, $0, 9
							slt $t2, $t0, $t1
							slt $t3, $t1, $t2`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(5)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(9)))
		Expect(*(sim.VM.GetReg("t2"))).To(Equal(int32(1)))
		Expect(*(sim.VM.GetReg("t3"))).To(Equal(int32(0)))
	})
})
