// Package to provides a set of useful functions for type conversion.
package to

import (
	"fmt"
	"strconv"
)

const (
	b   = "bool"
	f32 = "float32"
	f64 = "float64"
	i   = "int"
	i8  = "int8"
	i16 = "int16"
	i32 = "int32"
	i64 = "int64"
	s   = "string"
	u   = "uint"
	u8  = "uint8"
	u16 = "uint16"
	u32 = "uint32"
	u64 = "uint64"
)

// conversionError is a string format used for error messages when a
// conversion fails.
var conversionError = "unable to convert %v type of %T to %s"

// Float64 convert given 'v' to float64 type.
func Float32(v interface{}) (float32, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		return float32(v), nil
	case float64:
		return float32(v), nil
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case nil:
		return 0, nil
	case string:
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, f32)
		}
		return float32(f), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, f32)
	}
}

// Float64 convert given 'v' to float64 type.
func Float64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		return float64(v), nil
	case float64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case nil:
		return 0, nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, f64)
		}
		return float64(f), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, f64)
	}
}
