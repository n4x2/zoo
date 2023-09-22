// Package to provides a set of useful functions for type conversion.
package to

import (
	"errors"
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

// negativeValueError error returns for attempt convert value less
// than 0 for unsigned integers conversion.
var negativeValueError = errors.New("negative value is not allowed")

// Bool convert given 'v' to boolean type.
func Bool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case bool:
		return v, nil
	case float32:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case int:
		return v != 0, nil
	case int8:
		return v != 0, nil
	case int16:
		return v != 0, nil
	case int32:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case nil:
		return false, nil
	case string:
		return strconv.ParseBool(v)
	case uint:
		return v != 0, nil
	case uint8:
		return v != 0, nil
	case uint16:
		return v != 0, nil
	case uint32:
		return v != 0, nil
	case uint64:
		return v != 0, nil
	default:
		return false, fmt.Errorf(conversionError, v, v, b)
	}
}

// Float32 convert given 'v' to float32 type.
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

// Int convert given 'v' to int type.
func Int(v interface{}) (int, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case nil:
		return 0, nil
	case string:
		f, err := strconv.ParseInt(v, 0, 0)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, i)
		}
		return int(f), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, i)
	}
}

// Int8 convert given 'v' to int8 type.
func Int8(v interface{}) (int8, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		return int8(v), nil
	case float64:
		return int8(v), nil
	case int:
		return int8(v), nil
	case int8:
		return v, nil
	case int16:
		return int8(v), nil
	case int32:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case nil:
		return 0, nil
	case string:
		f, err := strconv.ParseInt(v, 0, 8)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, i8)
		}
		return int8(f), nil
	case uint:
		return int8(v), nil
	case uint8:
		return int8(v), nil
	case uint16:
		return int8(v), nil
	case uint32:
		return int8(v), nil
	case uint64:
		return int8(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, i8)
	}
}

// Int16 convert given 'v' to int16 type.
func Int16(v interface{}) (int16, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		return int16(v), nil
	case float64:
		return int16(v), nil
	case int:
		return int16(v), nil
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	case int32:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case nil:
		return 0, nil
	case string:
		f, err := strconv.ParseInt(v, 0, 16)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, i16)
		}
		return int16(f), nil
	case uint:
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case uint16:
		return int16(v), nil
	case uint32:
		return int16(v), nil
	case uint64:
		return int16(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, i16)
	}
}

// Int32 convert given 'v' to int32 type.
func Int32(v interface{}) (int32, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		return int32(v), nil
	case float64:
		return int32(v), nil
	case int:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		return int32(v), nil
	case nil:
		return 0, nil
	case string:
		f, err := strconv.ParseInt(v, 0, 32)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, i32)
		}
		return int32(f), nil
	case uint:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		return int32(v), nil
	case uint64:
		return int32(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, i32)
	}
}

// Int64 convert given 'v' to int64 type.
func Int64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case nil:
		return 0, nil
	case string:
		f, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, i64)
		}
		return f, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, i64)
	}
}

// Uint convert given 'v' to uint type.
func Uint(v interface{}) (uint, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint(v), nil
	case float64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint(v), nil
	case int:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint(v), nil
	case int8:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint(v), nil
	case int16:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint(v), nil
	case int32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint(v), nil
	case nil:
		return 0, nil
	case string:
		u, err := strconv.ParseInt(v, 0, 0)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, u)
		}

		if u < 0 {
			return 0, negativeValueError
		}
		return uint(u), nil
	case uint:
		return v, nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint64:
		return uint(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, i)
	}
}

// Uint8 convert given 'v' to uint8 type.
func Uint8(v interface{}) (uint8, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint8(v), nil
	case float64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint8(v), nil
	case int:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint8(v), nil
	case int16:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint8(v), nil
	case int32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint8(v), nil
	case int64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint8(v), nil
	case nil:
		return 0, nil
	case string:
		u, err := strconv.ParseInt(v, 0, 8)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, u8)
		}

		if u < 0 {
			return 0, negativeValueError
		}
		return uint8(u), nil
	case uint:
		return uint8(v), nil
	case uint8:
		return v, nil
	case uint16:
		return uint8(v), nil
	case uint32:
		return uint8(v), nil
	case uint64:
		return uint8(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, u8)
	}
}

// Uint16 convert given 'v' to unt16 type.
func Uint16(v interface{}) (uint16, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint16(v), nil
	case float64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint16(v), nil
	case int:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint16(v), nil
	case int32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint16(v), nil
	case int64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint16(v), nil
	case nil:
		return 0, nil
	case string:
		u, err := strconv.ParseInt(v, 0, 16)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, u16)
		}

		if u < 0 {
			return 0, negativeValueError
		}
		return uint16(u), nil
	case uint:
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint32:
		return uint16(v), nil
	case uint64:
		return uint16(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, u16)
	}
}

// Uint32 convert given 'v' to uint32 type.
func Uint32(v interface{}) (uint32, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint32(v), nil
	case float64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint32(v), nil
	case int:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint32(v), nil
	case nil:
		return 0, nil
	case string:
		u, err := strconv.ParseInt(v, 0, 32)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, u32)
		}

		if u < 0 {
			return 0, negativeValueError
		}
		return uint32(u), nil
	case uint:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		return uint32(v), nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, u32)
	}
}

// Uint64 convert given 'v' to uint64 type.
func Uint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case float32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint64(v), nil
	case int:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, negativeValueError
		}
		return uint64(v), nil
	case nil:
		return 0, nil
	case string:
		u, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			return 0, fmt.Errorf(conversionError, v, v, u64)
		}

		if u < 0 {
			return 0, negativeValueError
		}
		return uint64(u), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	default:
		return 0, fmt.Errorf(conversionError, v, v, u64)
	}
}

// String convert given 'v' to string type.
func String(v interface{}) (string, error) {
	switch v := v.(type) {
	case bool:
		return strconv.FormatBool(v), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case int:
		return strconv.Itoa(v), nil
	case int8:
		return strconv.Itoa(int(v)), nil
	case int16:
		return strconv.Itoa(int(v)), nil
	case int32:
		return strconv.Itoa(int(v)), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case nil:
		return "", nil
	case string:
		return v, nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	default:
		return "", fmt.Errorf(conversionError, v, v, s)
	}
}
