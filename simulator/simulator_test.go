package simulator_test

import (
	"time"

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

	It("performs SLTI OP", func() {
		sim = NewSimulator(`addi $t0, $0, 5
							slti $t1, $t0, 9
							slti $t2, $t0, 4`)
		sim.Run()
		Expect(*(sim.VM.GetReg("t0"))).To(Equal(int32(5)))
		Expect(*(sim.VM.GetReg("t1"))).To(Equal(int32(1)))
		Expect(*(sim.VM.GetReg("t2"))).To(Equal(int32(0)))
	})

	It("performs J op", func() {
		sim = NewSimulator(`
		main:
			j skip_addition
			addi $t0, $t0, 77

		skip_addition:
			addi $t0, $t0, 3
		`)

		sim.Run()
		Expect(*sim.VM.GetReg("t0")).To(Equal(int32(3)))
	})

	It("performs BEQ jumps", func(done Done) {
		sim = NewSimulator(`
		addi $t0, $0, 10 # set t0 to 10
		addi $t1, $0, 0 # set t1 to 0

		loop:
			beq $t1, $t0, end
			addi $t1, $t1, 1
			j loop

		end:
		`)

		// simulation runs in go-routine in case of infinite loop.
		// timeout after 200ms
		go sim.Run()

		time.Sleep(20 * time.Millisecond)
		Expect(*sim.VM.GetReg("t0")).To(Equal(int32(10)))
		close(done)
	}, 0.2)

	It("performs BNE jumps", func(done Done) {
		sim = NewSimulator(`
		addi $t0, $0, 0		# set t0 to 10
		addi $t1, $0, 10	# set t1 to 0

		loop:
			slt $t2, $t1, $t0	# t1 < t2
			bne $t2, $0, end
			addi $t1, $t1, -1
			j loop

		end:
		`)

		go sim.Run()

		time.Sleep(20 * time.Millisecond)
		Expect(*sim.VM.GetReg("t1")).To(Equal(int32(-1)))
		close(done)
	}, 0.2)
})
