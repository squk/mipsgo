package simulator_test

import (
	. "github.com/ctnieves/mipsgo/simulator"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("EdgeCases", func() {
	var sim Simulator

	It("handles mistyped registers", func() {
		sim = NewSimulator(`
			sll $ttt0, $s0, 4
			sll $randomstring, $s0, 4
		`)

		sim.Run()
	})

	It("handles mistyped operations", func() {
		sim = NewSimulator(`
			sllll $t0, $s0, 4
		`)

		sim.Run()
	})
})
