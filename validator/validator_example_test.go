package validator_test

import (
	"fmt"

	"github.com/n4x2/zoo/validator"
)

type User struct {
	Name string `json:"name" validate:"required"`
}

func Example() {
	var u User

	u.Name = ""

	v := validator.New()

	e, err := v.Validate(u)
	if err != nil {
		fmt.Print(err)
	}

	for _, err := range e {
		fmt.Println(err)
	}

	// Output:
	// {name [name is required]}
}

func Enum(f, v, p string) error { return nil }

func ExampleValidator_AddRule() {
	v := validator.New()

	err := v.AddRule("enum", Enum)
	if err != nil {
		fmt.Print(err)
	}

	// Output:
}

func ExampleValidator_SetValidatorTag() {
	v := validator.New()

	err := v.SetValidatorTag("v")
	if err != nil {
		fmt.Println(err)
	}

	// Output:
}
