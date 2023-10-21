package gogo

import (
	"fmt"
	"reflect"
	"runtime"
)

// 甭管是啥玩意，我给你打印出来🐱
func superPrintf(any interface{}) string {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			// print stack
			stack := make([]byte, 1024*8)
			stack = stack[:runtime.Stack(stack, false)]
			fmt.Println("super Print panic error, with err", panicErr, "with val", any, "with stack", string(stack))
		}
	}()
	if any == nil {
		return "nil"
	}
	v := reflect.ValueOf(any)
	t := v.Type()
	switch t.Kind() {
	case reflect.Invalid:
		return "nil"
	case reflect.Struct:
		return printStruct(v)
	case reflect.Slice:
		return printSlice(v)
	case reflect.Map:
		return printMap(v)
	case reflect.Ptr:
		elem := v.Elem()
		if elem.IsValid() {
			return superPrintf(elem.Interface())
		} else {
			return "nil"
		}
	default:
		return fmt.Sprintf("%v", any)
	}
}

func printStruct(v reflect.Value) string {
	t := v.Type()
	s := ""
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		s = fmt.Sprintf("%s%s: %s, ", s, f.Name, superPrintf(v.Field(i).Interface()))
	}
	return fmt.Sprintf("%s{%s}", t.Name(), s)

}

func printSlice(v reflect.Value) string {
	s := ""
	for i := 0; i < v.Len(); i++ {
		s = fmt.Sprintf("%s%s, ", s, superPrintf(v.Index(i).Interface()))
	}
	return fmt.Sprintf("[%s]", s)
}

func printMap(v reflect.Value) string {
	s := ""
	for _, k := range v.MapKeys() {
		s = fmt.Sprintf("%s%s: %s, ", s, superPrintf(k.Interface()), superPrintf(v.MapIndex(k).Interface()))
	}
	return fmt.Sprintf("{%s}", s)
}
