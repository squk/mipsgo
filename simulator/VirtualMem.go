package simulator

const VIRTUAL_MEMORY_SIZE int32 = (2 << 16)

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
func (m *VMem) SetWord(pos int, val WORD) {
	if pos > 0 && pos < int(VIRTUAL_MEMORY_SIZE) {
		m.Mem[pos] = val
	}
}

func (m *VMem) GetWord(pos int) WORD {
	if pos > 0 && pos < int(VIRTUAL_MEMORY_SIZE) {
		return m.Mem[pos]
	} else {
		return 0
	}
}

func (m *VMem) Wipe() {
	m.Mem = make([]WORD, VIRTUAL_MEMORY_SIZE)
}
