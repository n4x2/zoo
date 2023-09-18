// Package is provides functions for checking types and values.
package is

import (
	"reflect"

	"github.com/n4x2/zoo/constraints"
)

// Bool checks if the value is a boolean.
func Bool(v interface{}) bool {
	_, ok := v.(bool)
	return ok
}

// Byte checks if the value is a byte.
func Byte(v interface{}) bool {
	_, ok := v.(byte)
	return ok
}

// Contain checks if v is present in s.
func Contain[S ~[]E, E comparable](s S, v E) bool {
	for i := range s {
		if v == s[i] {
			return true
		}
	}
	return false
}

// ContainOneOf checks if part of ss is present in s.
func ContainOneOf[S comparable](s []S, ss []S) bool {
	for i := range ss {
		if Contain(s, ss[i]) {
			return true
		}
	}
	return false
}

// Equal checks if two comparable values are equal.
func Equal[T comparable](v, vv T) bool {
	return v == vv
}

// Error checks if the value is an error.
func Error(v interface{}) bool {
	_, ok := v.(error)
	return ok
}

// Float checks if the value is a float32 or float64.
func Float(v interface{}) bool {
	switch v.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}

// Int checks if the value is an integer.
func Int(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return true
	default:
		return false
	}
}

// Range checks if 'v' in range 'b' and 'e'.
func Range[T constraints.Number](b, e, v T) bool {
	return v >= b && v <= e
}

// Rune checks if the value is a rune.
func Rune(v interface{}) bool {
	_, ok := v.(rune)
	return ok
}

// Slice checks if the value is a slice.
func Slice(v interface{}) bool {
	s := reflect.ValueOf(v)
	return s.Kind() == reflect.Slice
}

// String checks if the value is a string.
func String(v interface{}) bool {
	_, ok := v.(string)
	return ok
}

// Struct checks if the value is a struct.
func Struct(v interface{}) bool {
	s := reflect.ValueOf(v)
	return s.Kind() == reflect.Struct
}

// Uint checks if the value is an unsigned integer.
func Uint(v interface{}) bool {
	switch v.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}
