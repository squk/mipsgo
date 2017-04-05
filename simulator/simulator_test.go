package simulator_test

import (
	. "github.com/ctnieves/mipsgo/simulator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simulator", func() {
	var sim Simulator

	It("performs addition", func() {
		sim = NewSimulator(`add $t0, $0, 5
							add $t1, $t0, 10`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(5)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(15)))
	})
})
