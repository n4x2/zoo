package to_test

import (
	"fmt"

	"github.com/n4x2/zoo/to"
)

func ExampleFloat32() {
	f, err := to.Float32(true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", f, f)
	// Output:
	// 1 type of float32
}

func ExampleFloat64() {
	f, err := to.Float64("3.14")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", f, f)
	// Output:
	// 3.14 type of float64
}

func ExampleBool() {
	b, err := to.Bool(0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", b, b)
	// Output:
	// false type of bool
}

func ExampleInt() {
	i, err := to.Int(-1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", i, i)
	// Output:
	// -1 type of int
}

func ExampleInt8() {
	i, err := to.Int8(-1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", i, i)
	// Output:
	// -1 type of int8
}

func ExampleInt16() {
	i, err := to.Int16(-1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", i, i)
	// Output:
	// -1 type of int16
}

func ExampleInt32() {
	i, err := to.Int32(-1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", i, i)
	// Output:
	// -1 type of int32
}

func ExampleInt64() {
	i, err := to.Int64(-1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v type of %T", i, i)
	// Output:
	// -1 type of int64
}
