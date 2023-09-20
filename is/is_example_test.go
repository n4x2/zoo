package is_test

import (
	"errors"
	"fmt"

	"github.com/n4x2/zoo/is"
)

func ExampleAlpha() {
	t := []string{"orange", "I like apple!"}
	for _, v := range t {
		fmt.Println(is.Alpha(v))
	}

	// Output:
	// true
	// false
}

func ExampleAlphaDash() {
	t := []string{"user_123", "command-line"}
	for _, v := range t {
		fmt.Println(is.AlphaDash(v))
	}

	// Output:
	// true
	// true
}

func ExampleAlphaNumeric() {
	t := []string{"user_123", "user123"}
	for _, v := range t {
		fmt.Println(is.AlphaNumeric(v))
	}

	// Output:
	// false
	// true
}

func ExampleASCII() {
	fmt.Println(is.ASCII("Hello World!"))

	// Output:
	// true
}

func ExampleBool() {
	s := []interface{}{
		false,
		123,
	}

	for _, v := range s {
		fmt.Println(is.Bool(v))
	}

	// Output:
	// true
	// false
}

func ExampleByte() {
	s := []interface{}{
		byte('a'),
		"a",
	}

	for _, v := range s {
		fmt.Println(is.Byte(v))
	}

	// Output:
	// true
	// false
}

func ExampleContain() {
	var s = []int{1, 2, 3, 100, 99}
	var a, b = 100, 88

	fmt.Println(is.Contain(s, a))
	fmt.Println(is.Contain(s, b))

	// Output:
	// true
	// false
}

func ExampleContainOneOf() {
	var s = []int{10, 20, 0, 100}
	var ss = []int{2, 20, 200}
	var a = []int{1, 2, 3}

	fmt.Println(is.ContainOneOf(s, ss))
	fmt.Println(is.ContainOneOf(s, a))

	// Output:
	// true
	// false
}

func ExampleEmail() {
	examples := []string{
		"user@example.com",
		"invalid_email@email",
		"another_user@example.com",
	}

	for _, v := range examples {
		fmt.Println(is.Email(v))
	}
	// Output:
	// true
	// false
	// true
}

func ExampleEqual() {
	var a, b = 1, 2
	var c, d = [3]byte{'a', 'b', 'c'}, [3]byte{'c', 'b', 'a'}
	var e, f = "equal", "equal"

	fmt.Println(is.Equal(a, b))
	fmt.Println(is.Equal(c, d))
	fmt.Println(is.Equal(e, f))

	// Output:
	// false
	// false
	// true
}

func ExampleError() {
	var e = errors.New("its error")
	fmt.Println(is.Error(e))

	// Output:
	// true
}

func ExampleFloat() {
	s := []interface{}{
		0,
		123,
		1000.1,
	}

	for _, v := range s {
		fmt.Println(is.Float(v))
	}

	// Output:
	// false
	// false
	// true
}

func ExampleGreaterThan() {
	fmt.Print(is.GreaterThan(1, 2))
	// Output:
	// false
}

func ExampleGreaterThanEqual() {
	fmt.Print(is.GreaterThanEqual(2, 2))
	// Output:
	// true
}

func ExampleInt() {
	s := []interface{}{
		rune('a'),
		0,
	}

	for _, v := range s {
		fmt.Println(is.Int(v))
	}

	// Output:
	// true
	// true
}

func ExampleLongitude() {
	examples := []string{"-179.999999", "180.000001", "invalid"}

	for _, v := range examples {
		fmt.Println(is.Longitude(v))
	}
	// Output:
	// true
	// false
	// false
}

func ExampleLessThan() {
	fmt.Print(is.LessThan(1, 2))
	// Output:
	// true
}

func ExampleLessThanEqual() {
	fmt.Print(is.LessThanEqual(2, 2))
	// Output:
	// true
}

func ExampleLatitude() {
	examples := []string{"-90.000001", "90.000001", "-89.999999", "invalid"}

	for _, v := range examples {
		fmt.Println(is.Latitude(v))
	}
	// Output:
	// false
	// false
	// true
	// false
}

func ExampleLowercase() {
	fmt.Println(is.Lowercase("cigarette"))
	// Output:
	// true
}

func ExampleNumber() {
	t := []string{"3", "3.14"}
	for _, v := range t {
		fmt.Println(is.Number(v))
	}

	// Output:
	// true
	// false
}

func ExampleNumeric() {
	t := []string{"3", "3.14"}
	for _, v := range t {
		fmt.Println(is.Numeric(v))
	}

	// Output:
	// true
	// true
}

func ExampleRange() {
	fmt.Println(is.Range('a', 'z', 'd'))
	fmt.Println(is.Range('a', 'z', 'H'))
	fmt.Println(is.Range(-100, 100, 2))
	fmt.Println(is.Range(-0.001, 1.1, 0.1))

	// Output:
	// true
	// false
	// true
	// true
}

func ExampleRune() {
	var r = rune('a')
	fmt.Println(is.Rune(r))

	// Output:
	// true
}

func ExampleSlice() {
	var s = []byte("abcde")
	var ss = []int{1, 2, 3, 10}

	fmt.Println(is.Slice(s))
	fmt.Println(is.Slice(ss))

	// Output:
	// true
	// true
}

func ExampleString() {
	var s = "this is string!"
	fmt.Println(is.String(s))

	// Output:
	// true
}

type Example struct {
	A, B string
}

func ExampleStruct() {
	var s Example
	fmt.Println(is.Struct(s))

	// Output:
	// true
}

func ExampleUint() {
	s := []interface{}{
		byte('a'),
		-1,
	}

	for _, v := range s {
		fmt.Println(is.Uint(v))
	}

	// Output:
	// true
	// false
}

func ExampleULID() {
	t := []string{
		"01AN4Z07BY79KA1307SR9X4MV3", // True ULID
		"not_a_valid_ulid",           // False ULID
	}

	for _, v := range t {
		fmt.Println(is.ULID(v))
	}
	// Output:
	// true
	// false
}

func ExampleUppercase() {
	fmt.Println(is.Uppercase("HELLO WORLD"))
	// Output:
	// true
}

func ExampleUUID() {
	t := []string{
		"550e8400-e29b-41d4-a716-446655440000",
		"not_a_valid_uuid",
	}

	for _, v := range t {
		fmt.Println(is.UUID(v))
	}
	// Output:
	// true
	// false
}
