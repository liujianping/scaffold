package models

import (
	"reflect"
	"regexp"

	"github.com/revel/revel"
)

var emailPattern = regexp.MustCompile("^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$")

func required() revel.Validator {
	return revel.Required{}
}

func min(size int) revel.Validator {
	return revel.MinSize{size}
}

func max(size int) revel.Validator {
	return revel.MaxSize{size}
}

func length(size int) revel.Validator {
	return revel.Length{size}
}

func email() revel.Validator {
	return revel.Match{emailPattern}
}

func match(exp string) revel.Validator {
	return revel.Match{regexp.MustCompile(exp)}
}

func IsZero(v reflect.Value) bool {
	valid := true
	switch v.Kind() {
	case reflect.String:
		valid = len(v.String()) != 0
	case reflect.Ptr, reflect.Interface:
		valid = !v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		valid = v.Len() != 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		valid = v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		valid = v.Float() != 0
	case reflect.Bool:
		valid = v.Bool()
	case reflect.Invalid:
		valid = false // always invalid
	}
	return !valid
}
