package is_test

import (
	"errors"
	"fmt"

	"github.com/n4x2/zoo/is"
)

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
