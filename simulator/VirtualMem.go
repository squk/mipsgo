package simulator

import "fmt"

const VIRTUAL_MEMORY_SIZE int32 = (2 << 0xD)

type HWORD int16
type WORD int32
type DWORD int64

type VMem struct {
	Mem []WORD
}

func InitMemory() VMem {
	var v VMem
	v.Mem = make([]WORD, VIRTUAL_MEMORY_SIZE)
	return v
}

// REAL ADDRESS SPACE
func (m *VMem) SetWord(pos int32, val WORD) {
	if pos > 0 && pos < VIRTUAL_MEMORY_SIZE {
		m.Mem[pos] = val
	}
}

func (m *VMem) GetWord(pos int32) WORD {
	if pos > 0 && pos < VIRTUAL_MEMORY_SIZE {
		return m.Mem[pos]
	} else {
		return 0
	}
}

func (m *VMem) Wipe() {
	m.Mem = make([]WORD, VIRTUAL_MEMORY_SIZE)
}

func (m *VMem) ToText() string {
	mem := ""

	for _, word := range m.Mem {
		mem += fmt.Sprintf("%X", word)
	}
	return mem
}
