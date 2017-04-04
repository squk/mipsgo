package simulator_test

import (
	. "github.com/ctnieves/mipsgo/simulator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lexer", func() {
	var sim Simulator

	BeforeEach(func() {
	})

	It("lexes numbers and identifiers", func() {
		sim = NewSimulator("add $t0, $t1, $t2")
		sim.Run()
		Expect(sim.Lexer.Tokens[0].Category).To(Equal(KEYWORD)) // 'add'
		Expect(sim.Lexer.Tokens[1].Category).To(Equal(SYMBOL))
		Expect(sim.Lexer.Tokens[2].Category).To(Equal(TEXT)) // 't0'
		Expect(sim.Lexer.Tokens[5].Category).To(Equal(TEXT)) // 't1'
		Expect(sim.Lexer.Tokens[8].Category).To(Equal(TEXT)) // 't2'
	})
})
