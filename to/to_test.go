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
		name     string
		input    interface{}
		expected bool
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
			assert[float32](t, Float32, test.name, "float32", test.input, test.expected)
		})
		t.Run("float64: test "+test.name, func(t *testing.T) {
			t.Parallel()
			assert[float64](t, Float64, test.name, "float64", test.input, test.expected)
		})
	}
}
