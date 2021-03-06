package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"reflect"
)

const (
	alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
)

func init() {
	// String
	gob.Register("")
	// Bool
	gob.Register(true)
	// Signed integer
	gob.Register(int(123))
	// Unsigned integer
	gob.Register(uint(123))
	// Float
	gob.Register(123.4)
	// Objects array
	gob.Register([]interface{}{})
	// Objects map indexed by string
	gob.Register(map[string]interface{}{})
	// Objects map indexed by integer
	gob.Register(map[int]interface{}{})
}

// Clone clones a plain object by gob.Encoder.Encode and gob.Decoder.Decode.
func Clone(obj interface{}) (interface{}, error) {
	if !reflect.ValueOf(obj).IsValid() {
		return nil, fmt.Errorf("Nil obj")
	}
	if reflect.ValueOf(obj).IsNil() {
		return obj, nil
	}
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	decoder := gob.NewDecoder(&buf)
	err := encoder.Encode(obj)
	if nil != err {
		return nil, err
	}
	cln := reflect.New(reflect.TypeOf(reflect.ValueOf(obj).Elem().Interface())).Interface()
	err = decoder.Decode(cln)
	if nil != err {
		return nil, err
	}
	return cln, nil
}

// RandString returns a random string at lenght n with characters [a-zA-Z0-9_].
// You should set the rand seed at your application initialization like:
// func init() {
// 	rand.Seed(int64(time.Now().Nanosecond()))
// }
func RandString(n int) string {
	if n == 0 {
		return ""
	}
	l := len(alpha)
	buff := bytes.Buffer{}
	for i := 0; i < n; i++ {
		buff.WriteRune(rune(alpha[rand.Intn(l)]))
	}
	return buff.String()
}
