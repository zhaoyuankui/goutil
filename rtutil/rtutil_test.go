package rtutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fun0(n int) string {
	return Caller(n)
}

func fun1(n int) string {
	return fun0(n)
}

func fun2(n int) string {
	return fun1(n)
}

func fun3(n int) string {
	return fun2(n)
}

type ss struct{}

func (s ss) fun(n int) string {
	return fun3(n)
}

func (s *ss) pfun(n int) string {
	return fun3(n)
}

func Test_Caller(t *testing.T) {
	assert.Equal(t, fun3(0), "fun0")
	assert.Equal(t, fun3(1), "fun1")
	assert.Equal(t, fun3(2), "fun2")
	assert.Equal(t, fun3(3), "fun3")
	s := ss{}
	assert.Equal(t, s.fun(4), "ss.fun")
	assert.Equal(t, (&s).fun(4), "ss.fun")
}
