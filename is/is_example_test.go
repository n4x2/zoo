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
