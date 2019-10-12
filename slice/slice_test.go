// test_util.go implements tests for util.go
// @Author zhaoyuankui@p1.com
// @Date 2018/12/04

package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestRetrieve struct {
	A string
}

func Test_RetrieveField(t *testing.T) {
	array := []TestRetrieve{
		{A: "aaa"},
		{A: "bbb"},
		{A: "ccc"},
	}
	res := SliceRetrieveField(array, "A")
	As := res.([]string)
	if !(As[0] == "aaa" && As[1] == "bbb" && As[2] == "ccc") {
		t.Error("Res", As)
	}
}

func Test_SliceAscertain(t *testing.T) {
	array := []interface{}{1, 2, 3, 4}
	var intArry []int
	intArry = SliceAscertain(array, 0).([]int)
	if intArry[0] != 1 {
		t.Error("Slice ascertain failed.")
	}
}

func Test_SliceBlur(t *testing.T) {
	array := []int{1, 2, 3, 4}
	var infArry []interface{}
	infArry = SliceBlur(array).([]interface{})
	if infArry[0].(int) != 1 {
		t.Error("Slice blur failed.")
	}
}

type testType struct {
	A int
	B int64
}

func Test_testCondition(t *testing.T) {
	slice := []testType{
		{
			A: 1,
			B: int64(2),
		},
		{
			A: 3,
			B: int64(4),
		},
	}
	field := reflect.ValueOf(slice[0].A)
	if !testCondition(field, "in", []int{1, 2, 3}) {
		t.Error("Test in failed")
	}
	if testCondition(field, "in", []int{2, 3, 4}) {
		t.Error("Test in failed")
	}

	if testCondition(field, "not in", []int{1, 2, 3}) {
		t.Error("Test not in failed")
	}
	if !testCondition(field, "not in", []int{2, 3, 4}) {
		t.Error("Test not in failed")
	}

	if !testCondition(field, "gt", 0) {
		t.Error("Test gt failed")
	}
	if testCondition(field, "gt", 1) {
		t.Error("Test gt failed")
	}
	if testCondition(field, "gt", 2) {
		t.Error("Test gt failed")
	}
	if !testCondition(field, "gt", int64(0)) {
		t.Error("Test gt failed")
	}

	if !testCondition(field, "eq", 1) {
		t.Error("Test eq failed")
	}
	if testCondition(field, "eq", 2) {
		t.Error("Test eq failed")
	}
	if !testCondition(field, "eq", int64(1)) {
		t.Error("Test eq failed")
	}
}

func Test_SliceFilterByField(t *testing.T) {
	slice := []testType{
		{
			A: 1,
			B: int64(2),
		},
		{
			A: 3,
			B: int64(4),
		},
	}
	res := SliceFilterByField(slice, "A", "eq", 1).([]testType)
	if len(res) != 1 || res[0].A != 1 {
		t.Error("Test SliceFilterByField failed")
	}
	res = SliceFilterByField(slice, "B", "eq", 2).([]testType)
	if len(res) != 1 || res[0].B != int64(2) {
		t.Error("Test SliceFilterByField failed")
	}
	res = SliceFilterByField(slice, "B", "eq", "abc").([]testType)
	if len(res) != 0 {
		t.Error("Test SliceFilterByField failed")
	}
	res2 := SliceFilterByField(123, "B", "eq", "abc")
	if nil != res2 {
		t.Error("Test SliceFilterByField failed")
	}
	res2 = SliceFilterByField(slice, "C", "eq", 2)
	if nil != res2 {
		t.Error("Test SliceFilterByField failed")
	}

	res = SliceFilterByField(slice, "A", "in", []int{1, 2, 3}).([]testType)
	if len(res) != 2 || res[0].A != 1 || res[1].A != 3 {
		t.Error("Test SliceFilterByField failed")
	}
	res = SliceFilterByField(slice, "B", "in", []int{1, 2, 3}).([]testType)
	if len(res) != 1 || res[0].A != 1 {
		t.Error("Test SliceFilterByField failed")
	}
}

func Test_SliceSubtract(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{2, 3, 4, 7}
	res := SliceSubtract(s1, s2).([]int)
	if !reflect.DeepEqual(res, []int{1, 5}) {
		t.Error("SliceSubtract failed")
	}
	res = SliceSubtract(s2, s1).([]int)
	if !reflect.DeepEqual(res, []int{7}) {
		t.Error("SliceSubtract failed")
	}
	res = SliceSubtract(s2, s2).([]int)
	if !reflect.DeepEqual(len(res), 0) {
		t.Error("SliceSubtract failed")
	}
	res = SliceSubtract(s2, []int{}).([]int)
	if !reflect.DeepEqual(res, s2) {
		t.Error("SliceSubtract failed")
	}
	res = SliceSubtract([]int{}, s2).([]int)
	if !reflect.DeepEqual(len(res), 0) {
		t.Error("SliceSubtract failed")
	}
}

func Test_InSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if !InSlice(s, 2) {
		t.Error("InSlice failed")
	}
	if InSlice(s, 6) {
		t.Error("InSlice failed")
	}
	if InSlice(s, "abc") {
		t.Error("InSlice failed")
	}
	if InSlice(123, "abc") {
		t.Error("InSlice failed")
	}
}

func Test_MakeSlice(t *testing.T) {
	initial := 1
	s := MakeSlice(5, initial).([]int)
	if len(s) != 5 {
		t.Error("Slice length error")
	}
	if s[0] != 1 {
		t.Error("Slice initial value error")
	}
}

type bar interface {
	Bar()
}

type foo interface {
	Foo()
}

type foo1 struct{}

func (f foo1) Foo() {}

func Test_SliceUpgrade(t *testing.T) {
	foo1s := []foo1{foo1{}, foo1{}, foo1{}}
	foos := SliceUpgrade(foo1s, reflect.TypeOf((*foo)(nil)).Elem()).([]foo)
	assert.Equal(t, len(foos), 3)
	// Upgrade type not valid
	assert.Panics(t, func() {
		SliceUpgrade(foo1s, reflect.TypeOf((*bar)(nil)).Elem())
	})
	// No elements
	foos = SliceUpgrade([]foo1{}, reflect.TypeOf((*foo)(nil)).Elem()).([]foo)
	assert.Equal(t, len(foos), 0)
	// Not slice type
	foos = SliceUpgrade(1, reflect.TypeOf((*foo)(nil)).Elem()).([]foo)
	assert.Equal(t, foos, []foo(nil))
	// Nil pointer
	foos = SliceUpgrade(nil, reflect.TypeOf((*foo)(nil)).Elem()).([]foo)
	assert.Equal(t, foos, []foo(nil))
}
