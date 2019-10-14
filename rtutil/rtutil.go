// Package rtutil implements runtime utilities.
package rtutil

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// RecoverError converts the recover value of panic to an error with the stack trace.
func RecoverError(r interface{}) error {
	buf := make([]byte, 1<<18)
	n := runtime.Stack(buf, false)
	return fmt.Errorf("Panic recovered. Reason: %v. Stack: %s", r, buf[0:n])
}

// Caller returns the simple name of the caller function that skips n stack frame.
func Caller(n int) string {
	// +1 skips Caller itself
	pc, _, _, ok := runtime.Caller(n + 1)
	if !ok {
		return ""
	}

	fnc := runtime.FuncForPC(pc)

	fullName := fnc.Name()
	idx := strings.LastIndex(fullName, "/")
	list := strings.Split(fullName[idx+1:], ".")
	name := list[len(list)-1]

	fmt.Printf("list: %+v\n", list)
	if len(list) > 2 {
		// method of struct
		sname := list[len(list)-2]
		if len(sname) > 1 && sname[0] == '(' && sname[len(sname)-1] == ')' {
			sname = sname[1 : len(sname)-2]
			if len(sname) > 0 && sname[0] == '*' {
				sname = sname[1:]
			}
		}
		fname := list[len(list)-1]
		name = sname + "." + fname
	}

	return name
}

// IsZero checks the value is the zero value of type t.
// Has the identical semantic of reflect.IsZero in go 1.13.
func IsZero(value reflect.Value, t reflect.Type) bool {
	if !value.IsValid() {
		return false
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(t).Interface())
}
