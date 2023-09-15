package pass_test

import (
	"fmt"

	"github.com/n4x2/zoo/pass"
)

func Example() {
	p, err := pass.Generate(true, true, 33)
	if err != nil {
		panic(err)
	}

	fmt.Print(p)
}
