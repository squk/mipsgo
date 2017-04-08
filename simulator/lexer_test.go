package simulator_test

import (
	"fmt"

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
		sim.PreProcess()
		Expect(sim.Lexer.Tokens[0].Category).To(Equal(KEYWORD)) // 'add'
		Expect(sim.Lexer.Tokens[1].Category).To(Equal(SYMBOL))
		Expect(sim.Lexer.Tokens[2].Category).To(Equal(TEXT)) // 't0'
		Expect(sim.Lexer.Tokens[5].Category).To(Equal(TEXT)) // 't1'
		Expect(sim.Lexer.Tokens[8].Category).To(Equal(TEXT)) // 't2'
	})

	It("skips comments", func() {
		sim = NewSimulator("# this is a test comment")
		sim.PreProcess()
		Expect(len(sim.Lexer.Tokens)).To(Equal(0))
	})

	It("identifies line numbers", func() {
		sim = NewSimulator(`
		add $t0, $t0, $t0
		sub $t0, $t0, $t0



		sub $t0, $t0, $t0




		add $t0, $t0, $t0
		`)
		sim.PreProcess()
		fmt.Println(sim.Lexer.GetTokens())
		Expect(sim.Parser.Instructions[0].LineNumber).To(Equal(2))
		Expect(sim.Parser.Instructions[1].LineNumber).To(Equal(3))
		Expect(sim.Parser.Instructions[2].LineNumber).To(Equal(7))
		Expect(sim.Parser.Instructions[3].LineNumber).To(Equal(12))
	})
})
