package alg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice(t *testing.T) {
	s := Slice(TElemInt)
	assert.NotNil(t, s)
	assert.IsType(t, ([]int)(nil), s)
	assert.Empty(t, s)
	s = Slice(TElemInt, 1, 2, 3)
	assert.Equal(t, s, []int{1, 2, 3})
	s = Slice(TElemString, "1", "2", "3")
	assert.Equal(t, s, []string{"1", "2", "3"})
}

func Test_IntSlice(t *testing.T) {
	s := IntSlice(1, 2, 3)
	assert.Equal(t, s, []int{1, 2, 3})
}

func Test_StringSlice(t *testing.T) {
	s := StringSlice("1", "2", "3")
	assert.Equal(t, s, []string{"1", "2", "3"})
}

func Test_FloatSlice(t *testing.T) {
	s := FloatSlice(0.1, 0.2, 0.3)
	assert.Equal(t, s, []float64{0.1, 0.2, 0.3})
}
