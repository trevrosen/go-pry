package pry

import (
	"errors"
	"fmt"
	"reflect"
)

// A runtime replacement for the append function
func Append(arr, elem interface{}) (interface{}, error) {
	if reflect.TypeOf(arr) != reflect.SliceOf(reflect.TypeOf(elem)) {
		return nil, errors.New(fmt.Sprintf("%T cannot append to %T.", elem, arr))
	}
	arrVal := reflect.ValueOf(arr)
	elemVal := reflect.ValueOf(elem)
	return reflect.Append(arrVal, elemVal).Interface(), nil
}

// A runtime replacement for the make function
func Make(t interface{}, args ...interface{}) (interface{}, error) {
	typ, isType := t.(reflect.Type)
	if !isType {
		return nil, errors.New(fmt.Sprintf("Invalid type %#v", t))
	}
	switch typ.Kind() {
	case reflect.Slice:
		if len(args) < 1 || len(args) > 2 {
			return nil, errors.New("Invalid number of arguments. Missing len or extra?")
		}
		length, isInt := args[0].(int)
		if !isInt {
			return nil, errors.New("Len is not int.")
		}
		capacity := length
		if len(args) == 2 {
			capacity, isInt = args[0].(int)
			if !isInt {
				return nil, errors.New("Len is not int.")
			}
		}
		slice := reflect.MakeSlice(typ, length, capacity)
		return slice.Interface(), nil

	case reflect.Chan:
		if len(args) > 1 {
			return nil, errors.New("Invalid number of arguments. Too many.")
		}
		size := 0
		if len(args) == 1 {
			var isInt bool
			size, isInt = args[0].(int)
			if !isInt {
				return nil, errors.New("Size is not int.")
			}
		}
		buffer := reflect.MakeChan(typ, size)
		return buffer.Interface(), nil

	default:
		return nil, errors.New(fmt.Sprintf("Unknown kind type %T", t))
	}
}
