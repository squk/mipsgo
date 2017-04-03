package simulator

import (
	"fmt"
	"io/ioutil"
)

type Simulator struct {
	Filename string
	Source   string
	VM       VirtualMachine
}

func ReadSource(filename string) Simulator {
	var s Simulator
	s.Filename = filename
	s.GetSource()

	return s
}

func (s *Simulator) GetSource() {
	b, err := ioutil.ReadFile(s.Filename)

	if err != nil {
		fmt.Print(err)
	}

	s.Source = b
}
