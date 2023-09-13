// Package is provides functions for checking types and values.
package is

import "reflect"

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
