package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type foo struct {
	A  string
	M  map[string]int
	M2 map[string]int
	C  chan int
	F  func() string
	a  string
}

func Test_Clone(t *testing.T) {
	f := &foo{
		A: "aaa",
		M: map[string]int{
			"one": 1,
			"two": 2,
		},
		C: make(chan int),
		F: func() string { return "Hello" },
		a: "aa",
	}
	c, _ := Clone(f)
	fc, ok := c.(*foo)
	if !ok {
		t.Error("Expect *foo type of fc")
	}
	assert.Equal(t, fc.A, "aaa")
	assert.Equal(t, fc.M, map[string]int{"one": 1, "two": 2})
	assert.Nil(t, fc.M2)
	assert.Nil(t, fc.C)
	assert.Nil(t, fc.F)
	assert.Empty(t, fc.a)

	// Nil value
	_, err := Clone(nil)
	assert.Contains(t, err.Error(), "Nil obj")

	// Nil pointer
	var nilP *int
	p, err := Clone(nilP)
	assert.Nil(t, err)
	intP, ok := p.(*int)
	assert.Nil(t, intP)
}
