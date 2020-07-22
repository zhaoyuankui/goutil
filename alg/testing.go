package alg

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertEqual tests whether the file contents equals to the object o.
func AssertEqual(t *testing.T, file string, o interface{}) {
	f, err := os.Open(file)
	assert.Nil(t, err)
	contents := make([]byte, 1024*1024)
	n, err := f.Read(contents)
	assert.Nil(t, err)

	var buf bytes.Buffer
	switch o.(type) {
	case []string:
		for _, s := range o.([]string) {
			buf.WriteString(s)
			buf.WriteString("\n")
		}
		assert.Equal(t, buf.Bytes(), contents[:n])
	case []int:
		for _, i := range o.([]int) {
			intBytes := int2Bytes(i)
			buf.Write(intBytes)
			buf.WriteString("\n")
		}
		assert.Equal(t, buf.String(), string(contents[:n]))
	case [][]int:
		for _, a := range o.([][]int) {
			for j, i := range a {
				intBytes := int2Bytes(i)
				buf.Write(intBytes)
				if j == len(a)-1 {
					buf.WriteString("\n")
				} else {
					buf.WriteString(" ")
				}
			}
		}
		assert.Equal(t, buf.String(), string(contents[:n]))
	default:
		panic(fmt.Sprintf("Type %s not supported.", reflect.TypeOf(o)))
	}
}

func int2Bytes(i int) []byte {
	intBytes := []byte{}
	for i > 0 {
		intBytes = append([]byte{byte('0' + i%10)}, intBytes...)
		i /= 10
	}
	return intBytes
}
