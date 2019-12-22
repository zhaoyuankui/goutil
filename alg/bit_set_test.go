package alg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBitSet(t *testing.T) {
	s := NewBitSet(0)
	assert.Equal(t, []byte(nil), s.bits)
	assert.Equal(t, uint(0), s.l)

	s = NewBitSet(1)
	assert.Equal(t, uint(1), s.l)
	assert.Equal(t, []byte{0}, s.bits)
	s.SetN(0)
	assert.Equal(t, []byte{1}, s.bits)
	assert.True(t, s.TestN(0))
	s.UnsetN(0)
	assert.Equal(t, []byte{0}, s.bits)
	assert.False(t, s.TestN(0))
	s.Set()
	assert.Equal(t, []byte{0xff}, s.bits)
	assert.True(t, s.TestN(0))
	s.Unset()
	assert.Equal(t, []byte{0}, s.bits)
	assert.False(t, s.TestN(0))

	s = NewBitSet(8)
	assert.Equal(t, []byte{0}, s.bits)

	s = NewBitSet(9)
	assert.Equal(t, []byte{0, 0}, s.bits)
	assert.Panics(t, func() { s.TestN(9) })
	assert.Panics(t, func() { s.SetN(9) })
	assert.Panics(t, func() { s.UnsetN(9) })
	assert.False(t, s.TestN(8))
	s.SetN(8)
	assert.True(t, s.TestN(8))
	assert.Equal(t, []byte{0, 1}, s.bits)

	s = NewBitSet(32)
	assert.Equal(t, []byte{0, 0, 0, 0}, s.bits)
	s = NewBitSet(36)
	assert.Equal(t, []byte{0, 0, 0, 0, 0}, s.bits)
	s.SetN(31)
	assert.True(t, s.TestN(31))
	s.UnsetN(31)
	assert.False(t, s.TestN(31))
	s.SetN(31)
	s.SetN(35)
	assert.Equal(t, []byte{0, 0, 0, 128, 8}, s.bits)
}
