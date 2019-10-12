// slice.go implements some utilities of slice.
// @Author zhaoyuankui@p1.com
// @Date 2018/12/04

package slice

import (
	"reflect"

	"github.com/zhaoyuankui/goutil/cmp"
)

// SliceRetrieveField retrieves values of a field of objects in a slice
// to another slice of the field type wrapped by reflect.Value.
func SliceRetrieveField(slice interface{}, field string) interface{} {
	kind := reflect.TypeOf(slice).Kind()
	if kind != reflect.Slice {
		return reflect.ValueOf(nil)
	}
	s := reflect.ValueOf(slice)
	l := s.Len()
	if l <= 0 {
		return reflect.ValueOf(nil)
	}
	fieldType := s.Index(0).FieldByName(field).Type()
	res := reflect.MakeSlice(reflect.SliceOf(fieldType), 0, l)
	for i := 0; i < l; i++ {
		obj := s.Index(i)
		val := obj.FieldByName(field)
		res = reflect.Append(res, val)
	}
	return res.Interface()
}

func SliceAscertain(slice interface{}, hint interface{}) interface{} {
	kind := reflect.TypeOf(slice).Kind()
	if kind != reflect.Slice {
		return reflect.ValueOf(nil)
	}
	elemType := reflect.TypeOf(hint)
	sliceType := reflect.SliceOf(elemType)
	s := reflect.ValueOf(slice)
	l := s.Len()
	res := reflect.MakeSlice(sliceType, 0, l)
	for i := 0; i < l; i++ {
		val := reflect.ValueOf(s.Index(i).Interface())
		// Try convert if types not match.
		if val.Type().Kind() != elemType.Kind() {
			res = reflect.Append(res, val.Convert(elemType))
		} else {
			res = reflect.Append(res, val)
		}
	}
	return res.Interface()
}

func SliceBlur(slice interface{}) []interface{} {
	kind := reflect.TypeOf(slice).Kind()
	if kind != reflect.Slice {
		return []interface{}(nil)
	}
	s := reflect.ValueOf(slice)
	l := s.Len()
	res := make([]interface{}, 0, l)
	for i := 0; i < l; i++ {
		res = append(res, s.Index(i).Interface())
	}
	return res
}

func SliceUpgrade(slice interface{}, t reflect.Type) interface{} {
	sliceType := reflect.SliceOf(t)
	if nil == slice {
		return reflect.Zero(sliceType).Interface()
	}
	kind := reflect.TypeOf(slice).Kind()
	if kind != reflect.Slice {
		return reflect.Zero(sliceType).Interface()
	}
	s := reflect.ValueOf(slice)
	l := s.Len()
	res := reflect.MakeSlice(sliceType, 0, l)
	for i := 0; i < l; i++ {
		res = reflect.Append(res, s.Index(i))
	}
	return res.Interface()
}

func SliceFilterByField(slice interface{}, field string, cond string, value interface{}) interface{} {
	kind := reflect.TypeOf(slice).Kind()
	if kind != reflect.Slice {
		return nil
	}
	s := reflect.ValueOf(slice)
	l := s.Len()
	t := reflect.TypeOf(slice)
	filteredSlice := reflect.MakeSlice(t, 0, 0)
	for i := 0; i < l; i++ {
		obj := s.Index(i)
		f := obj.FieldByName(field)
		if !f.IsValid() {
			return nil
		}
		if !testCondition(f, cond, value) {
			continue
		}
		filteredSlice = reflect.Append(filteredSlice, obj)
	}
	return filteredSlice.Interface()
}

func testCondition(f reflect.Value, cond string, value interface{}) bool {
	if cond == "in" || cond == "not in" {
		kind := reflect.TypeOf(value).Kind()
		if kind != reflect.Slice {
			return false
		}
	}
	switch cond {
	case "in":
		slice := reflect.ValueOf(value)
		for i, l := 0, slice.Len(); i < l; i = i + 1 {
			if isEq, err := cmp.Eq(slice.Index(i), f); isEq && nil == err {
				return true
			}
		}
		return false
	case "gt":
		if isGt, err := cmp.Gt(f, reflect.ValueOf(value)); isGt && nil == err {
			return true
		}
		return false
	case "eq":
		if isEq, err := cmp.Eq(f, reflect.ValueOf(value)); isEq && nil == err {
			return true
		}
		return false
	case "not in":
		return !testCondition(f, "in", value)
	case "lt":
		return !testCondition(f, "gt", value) && !testCondition(f, "eq", value)
	case "le":
		return !testCondition(f, "gt", value)
	case "ge":
		return testCondition(f, "gt", value) || testCondition(f, "eq", value)
	// Not support.
	default:
		return false
	}
}

func SliceSubtract(s1 interface{}, s2 interface{}) interface{} {
	kind1 := reflect.TypeOf(s1).Kind()
	kind2 := reflect.TypeOf(s2).Kind()
	if kind1 != reflect.Slice || kind1 != kind2 {
		return nil
	}
	slice1 := reflect.ValueOf(s1)
	slice2 := reflect.ValueOf(s2)
	boolType := reflect.TypeOf(true)
	boolTrue := reflect.ValueOf(true)
	elemType := reflect.TypeOf(s2).Elem()
	set := reflect.MakeMap(reflect.MapOf(elemType, boolType))
	for i, l := 0, slice2.Len(); i < l; i = i + 1 {
		value := slice2.Index(i)
		set.SetMapIndex(value, boolTrue)
	}
	subtractSlice := reflect.MakeSlice(reflect.TypeOf(s1), 0, 0)
	for i, l := 0, slice1.Len(); i < l; i = i + 1 {
		value := slice1.Index(i)
		if set.MapIndex(value).IsValid() {
			continue
		}
		subtractSlice = reflect.Append(subtractSlice, value)
	}
	return subtractSlice.Interface()
}

func InSlice(slice interface{}, value interface{}) bool {
	kind := reflect.TypeOf(slice).Kind()
	if kind != reflect.Slice {
		return false
	}
	s := reflect.ValueOf(slice)
	l := s.Len()
	for i := 0; i < l; i++ {
		if isEq, err := cmp.Eq(s.Index(i), reflect.ValueOf(value)); isEq && nil == err {
			return true
		}
	}
	return false
}

func MakePair(v1 interface{}, v2 interface{}) Pair {
	return Pair{v1, v2}
}

func MakeTriple(v1 interface{}, v2 interface{}, v3 interface{}) Triple {
	return Triple{v1, v2, v3}
}

func MakeSlice(l int, initial interface{}) interface{} {
	s := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(initial)), 0, l)
	for i := 0; i < l; i = i + 1 {
		s = reflect.Append(s, reflect.ValueOf(initial))
	}
	return s.Interface()
}
