package to

import (
	"fmt"
	"testing"
)

func assert[T any](t *testing.T, fn func(v any) (T, error), n, x string, v any, e bool) {
	t.Helper()
	t.Run(n, func(t *testing.T) {
		value, err := fn(v)
		if err == nil && e {
			t.Errorf("unexpected error %v", err)
		}

		if x != fmt.Sprintf("%T", value) && !e {
			t.Errorf("unexpected type %T", value)
		}
	})
}

func TestFloatingPoint(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name  string
		input interface{}
		err   bool
	}{
		{"bool true", true, false},
		{"bool false", false, false},
		{"float32", float32(0.1210), false},
		{"float64", float64(-10.1), false},
		{"int", int(-11), false},
		{"int8", int8(8), false},
		{"int16", int16(16), false},
		{"int32", int32(32), false},
		{"int64", int64(46), false},
		{"nil", nil, false},
		{"string 1.1", "1.1", false},
		{"string 200", "200", false},
		{"uint", uint(0), false},
		{"uint8", uint8(8), false},
		{"uint16", uint16(16), false},
		{"uint32", uint32(32), false},
		{"uint64", uint64(46), false},
		{"fail string ;", ";", true},
		{"fail slice []int", []int{0}, true},
	}
	for _, test := range tests {
		test := test
		t.Run("float32: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[float32](t, Float32, test.name, "float32", test.input, test.err)
		})
		t.Run("float64: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[float64](t, Float64, test.name, "float64", test.input, test.err)
		})
	}
}

func TestBool(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name  string
		input interface{}
		err   bool
	}{
		{"bool true", true, false},
		{"bool false", false, false},
		{"float32 0", float32(0), false},
		{"float64 0", float64(0), false},
		{"int 0", int(0), false},
		{"int8 0", int8(0), false},
		{"int16 0", int16(0), false},
		{"int32 0", int32(0), false},
		{"int64 0", int64(0), false},
		{"uint 0", uint(0), false},
		{"uint8 0", uint8(0), false},
		{"uint16 0", uint16(0), false},
		{"uint32 0", uint32(0), false},
		{"uint64 0", uint64(0), false},
		{"float32 1", float32(1), false},
		{"float64 1", float64(1), false},
		{"int 1", int(1), false},
		{"int8 1", int8(1), false},
		{"int16 1", int16(1), false},
		{"int32 1", int32(1), false},
		{"int64 1", int64(1), false},
		{"uint 1", uint(1), false},
		{"uint8 1", uint8(1), false},
		{"uint16 1", uint16(1), false},
		{"uint32 1", uint32(1), false},
		{"uint64 1", uint64(1), false},
		{"nil", nil, false},
		{"string 1", "1", false},
		{"string 0", "0", false},
		{"string t", "t", false},
		{"string T", "T", false},
		{"string true", "true", false},
		{"string True", "True", false},
		{"string TRUE", "TRUE", false},
		{"string f", "f", false},
		{"string F", "F", false},
		{"string false", "false", false},
		{"string False", "False", false},
		{"string FALSE", "FALSE", false},

		// fail
		{"string test", "test", true},
		{"slice", []int{0, 1}, true},
	}

	for _, test := range tests {
		test := test
		t.Run("bool: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[bool](t, Bool, test.name, "bool", test.input, test.err)
		})
	}
}
