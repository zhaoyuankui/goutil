// This source implements a bit set with byte slice.

package alg

import "fmt"

// BitSet defines the type of a bit set,
// where bits is the byte slice saves all bits,
// and l is the real lenth of the bit set.
type BitSet struct {
	bits []byte
	l    uint
}

// NewBitSet creates a bit set with lenth l and all bits unset.
// If l == 0, the byte slice would be nil, and the bit set has nonsense.
func NewBitSet(l uint) *BitSet {
	if l == 0 {
		return &BitSet{}
	}
	c := (l + 7) >> 3
	bitSet := &BitSet{
		bits: make([]byte, c, c),
		l:    l,
	}
	return bitSet
}

// Set sets all bits to 1.
func (s *BitSet) Set() {
	for i := 0; i < len(s.bits); i++ {
		s.bits[i] = 0xff
	}
}

// Unset sets all bits to 0.
func (s *BitSet) Unset() {
	for i := 0; i < len(s.bits); i++ {
		s.bits[i] = 0
	}
}

// SetN sets the nth bit to 1, where n ranges [0, s.l-1].
// If n out of the range or s.l == 0, it panics.
func (s *BitSet) SetN(n uint) {
	if n >= s.l {
		panic(fmt.Sprintf("SetN out of range. n: %d, s.l: %d", n, s.l))
	}
	s.bits[n>>3] |= masks[n&7]
}

// UnsetN sets the nth bit to 0, where n ranges [0, s.l-1].
// If n out of the range or s.l == 0, it panics.
func (s *BitSet) UnsetN(n uint) {
	if n >= s.l {
		panic(fmt.Sprintf("UnsetN out of range. n: %d, s.l: %d", n, s.l))
	}
	s.bits[n>>3] &= ^masks[n&7]
}

// TestN tests whether the nth bit is set.
// If set, it returns true, or else false.
func (s *BitSet) TestN(n uint) bool {
	if n >= s.l {
		panic(fmt.Sprintf("TestN out of range. n: %d, s.l: %d", n, s.l))
	}
	return s.bits[n>>3]&masks[n&7] > 0
}

var masks = [8]byte{
	0x01,
	0x02,
	0x04,
	0x08,
	0x10,
	0x20,
	0x40,
	0x80,
}
