package simulator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSimulator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Simulator Suite")
}
