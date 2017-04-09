package simulator

import (
	"fmt"
	"strconv"
	"strings"
)

// 4.096 KB of RAM, 1024 words
const VIRTUAL_MEMORY_SIZE int32 = (2 << 9)

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
		mem += fmt.Sprintf("%08X", uint32(word))
	}
	return mem
}

func (m *VMem) Write(hex string) {
	for i, s := range strings.Split(hex, " ") {
		if i < len(m.Mem) {
			if num, err := strconv.ParseUint(s, 16, 32); err == nil {
				m.Mem[i] = WORD(num)
			} else {
				fmt.Println(err)
				m.Mem[i] = 0
			}
		} else {
			break
		}
	}
}
