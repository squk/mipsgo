package simulator_test

import (
	. "github.com/ctnieves/mipsgo/simulator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simulator", func() {
	var sim Simulator

	It("performs ADD op", func() {
		sim = NewSimulator(`add $t0, $0, 5
							add $t1, $t0, 10`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(5)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(15)))
	})

	It("performs SUB op", func() {
		sim = NewSimulator(`add $t0, $0, 15
							sub $t1, $t0, 10`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(15)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(5)))
	})

	It("performs SLT OP", func() {
		sim = NewSimulator(`add $t0, $0, 5
							add $t1, $0, 9
							slt $t2, $t0, $t1
							slt $t3, $t1, $t2`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(5)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(9)))
		Expect(*(sim.VM.GetReg("t2"))).To(Equal(int32(1)))
		Expect(*(sim.VM.GetReg("t3"))).To(Equal(int32(0)))
	})
})
