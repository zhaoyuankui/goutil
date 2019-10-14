package rtutil

import (
	"reflect"
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

func Test_IsZero(t *testing.T) {
	var value reflect.Value
	i := 2
	assert.False(t, IsZero(value, reflect.TypeOf(i)))
	assert.False(t, IsZero(reflect.ValueOf(i), reflect.TypeOf(i)))
	i = 0
	assert.True(t, IsZero(reflect.ValueOf(i), reflect.TypeOf(i)))
	var s []int
	assert.True(t, IsZero(reflect.ValueOf(s), reflect.TypeOf(s)))
	s = []int{}
	assert.False(t, IsZero(reflect.ValueOf(s), reflect.TypeOf(s)))
	s = []int{1}
	assert.False(t, IsZero(reflect.ValueOf(s), reflect.TypeOf(s)))
	s = nil
	assert.True(t, IsZero(reflect.ValueOf(s), reflect.TypeOf(s)))
	assert.False(t, IsZero(value, reflect.TypeOf(s)))
}
