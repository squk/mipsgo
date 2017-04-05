package simulator_test

import (
	. "github.com/ctnieves/mipsgo/simulator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var sim Simulator

	It("parses insstructions and arguments", func() {
		sim = NewSimulator(`sub $s0, $s1, $s2
							add $t0, $t1, $t2
							sll $t1 4`)
		sim.Run()
		Expect(sim.Parser.Instructions[0]).To(Equal(Instruction{OpCode: 39, RD: 16, RS: 17, RT: 18}))
		Expect(sim.Parser.Instructions[1]).To(Equal(Instruction{OpCode: 1, RD: 8, RS: 9, RT: 10}))
		Expect(sim.Parser.Instructions[2]).To(Equal(Instruction{OpCode: 30, RD: 9, RS: -1, RT: -1, Immediate: 4}))
	})
})
