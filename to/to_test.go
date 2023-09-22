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
			t.Errorf("unexpected error %v", err.Error())
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

func TestSignedInteger(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name  string
		input interface{}
		err   bool
	}{
		{"bool true", true, false},
		{"bool false", false, false},
		{"float32", float32(-1), false},
		{"float64", float64(-1), false},
		{"int", int(-1), false},
		{"int8", int8(-1), false},
		{"int16", int16(-1), false},
		{"int32", int32(-1), false},
		{"int64", int64(-1), false},
		{"nil", nil, false},
		{"string -1", "-1", false},
		{"string 1", "1", false},
		{"uint", uint(1), false},
		{"uint8", uint8(1), false},
		{"uint16", uint16(1), false},
		{"uint32", uint32(1), false},
		{"uint64", uint64(1), false},
		{"fail string ;", ";", true},
		{"fail slice []int", []int{0}, true},
	}
	for _, test := range tests {
		test := test
		t.Run("int: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[int](t, Int, test.name, "int", test.input, test.err)
		})
		t.Run("int8: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[int8](t, Int8, test.name, "int8", test.input, test.err)
		})
		t.Run("int16: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[int16](t, Int16, test.name, "int16", test.input, test.err)
		})
		t.Run("int32: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[int32](t, Int32, test.name, "int32", test.input, test.err)
		})
		t.Run("int64: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[int64](t, Int64, test.name, "int64", test.input, test.err)
		})
	}
}

func TestUnsignedInteger(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name  string
		input interface{}
		err   bool
	}{
		{"bool true", true, false},
		{"bool false", false, false},
		{"float32", float32(1), false},
		{"float64", float64(1), false},
		{"int", int(1), false},
		{"int8", int8(1), false},
		{"int16", int16(1), false},
		{"int32", int32(1), false},
		{"int64", int64(1), false},
		{"nil", nil, false},
		{"string 1", "1", false},
		{"uint", uint(1), false},
		{"uint8", uint8(1), false},
		{"uint16", uint16(1), false},
		{"uint32", uint32(1), false},
		{"uint64", uint64(1), false},

		//fail
		{"fail string ;", ";", true},
		{"fail slice []int", []int{0}, true},
		{"fail float32", float32(-1), true},
		{"fail float64", float64(-1), true},
		{"fail int", int(-1), true},
		{"fail int8", int8(-1), true},
		{"fail int16", int16(-1), true},
		{"fail int32", int32(-1), true},
		{"fail int64", int64(-1), true},
		{"fail string -1", "-1", true},
	}
	for _, test := range tests {
		test := test
		t.Run("uint: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[uint](t, Uint, test.name, "uint", test.input, test.err)
		})
		t.Run("uint8: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[uint8](t, Uint8, test.name, "uint8", test.input, test.err)
		})
		t.Run("uint16: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[uint16](t, Uint16, test.name, "uint16", test.input, test.err)
		})
		t.Run("uint32: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[uint32](t, Uint32, test.name, "uint32", test.input, test.err)
		})
		t.Run("uint64: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[uint64](t, Uint64, test.name, "uint64", test.input, test.err)
		})
	}
}
