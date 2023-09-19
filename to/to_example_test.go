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
