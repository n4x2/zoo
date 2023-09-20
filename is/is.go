// Package is provides functions for checking types and values.
package is

import (
	"reflect"
	"strings"

	"github.com/n4x2/zoo/constraints"
	"github.com/n4x2/zoo/regex"
)

// Alpha checks if the value is letters a-z or A-Z.
func Alpha(v string) bool {
	return regex.Alpha.MatchString(v)
}

// AlphaDash checks if the value is letters, numbers, dash, and
// underscore.
func AlphaDash(v string) bool {
	return regex.AlphaDash.MatchString(v)
}

// AlphaNumeric checks if the value is letters and numbers.
func AlphaNumeric(v string) bool {
	return regex.AlphaNumeric.MatchString(v)
}

// ASCII checks if the value is ASCII characters.
func ASCII(v string) bool {
	return regex.ASCII.MatchString(v)
}

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

// Email checks if the value is valid email.
func Email(v string) bool {
	return regex.Email.MatchString(v)
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

// GreaterThan checks if 'v' is greater than 'p'.
func GreaterThan[T constraints.Number](v, p T) bool {
	return v > p
}

// GreaterThanEqual checks if 'v' is greater than or equal 'p'.
func GreaterThanEqual[T constraints.Number](v, p T) bool {
	return v >= p
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

// Latitude checks if the is valid latitude.
func Latitude(v string) bool {
	return regex.Latitude.MatchString(v)
}

// LessThan checks if 'v' is less than 'p'.
func LessThan[T constraints.Number](v, p T) bool {
	return v < p
}

// LessThanEqual checks if 'v' is less than or equal 'p'.
func LessThanEqual[T constraints.Number](v, p T) bool {
	return v <= p
}

// Longitude checks if the value is valid longitude.
func Longitude(v string) bool {
	return regex.Longitude.MatchString(v)
}

// Lowercase checks if the value is lower case.
func Lowercase(v string) bool {
	return strings.ToLower(v) == v
}

// Number checks if the value is numbers.
func Number(v string) bool {
	return regex.Number.MatchString(v)
}

// Numeric checks if the value is numeric type: integer or
// floating-point.
func Numeric(v string) bool {
	return regex.Numeric.MatchString(v)
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

// ULID checks if the value is valid ULID.
func ULID(v string) bool {
	return regex.ULID.MatchString(v)
}

// Uppercase checks if the value is upper case.
func Uppercase(v string) bool {
	return strings.ToUpper(v) == v
}

// UUID checks if the value is valid UUID.
func UUID(v string) bool {
	return regex.UUID.MatchString(v)
}
