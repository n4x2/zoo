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

type newRule struct{}

func (r *newRule) Name() string                  { return "newrule" }
func (r *newRule) Validate(f, v, p string) error { return nil }

func ExampleValidator_AddRule() {
	v := validator.New()

	err := v.AddRule(&newRule{})
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
