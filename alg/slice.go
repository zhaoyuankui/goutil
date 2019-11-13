package alg

import "reflect"

func Slice(hint ElemType, args ...interface{}) interface{} {
	elemType := reflect.TypeOf(hint)
	sliceType := reflect.SliceOf(elemType)
	slice := reflect.MakeSlice(sliceType, 0, len(args))
	for _, v := range args {
		slice = reflect.Append(slice, reflect.ValueOf(v))
	}
	return slice.Interface()
}
