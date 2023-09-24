package validator_test

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/n4x2/zoo/validator"
)

type Student struct {
	FName  string `json:"first_name" v:"lowercase"`
	MName  string `json:"middle_name" v:"-|lowercase"`
	LName  string `json:"last_name" v:"lowercase"`
	Age    int8   `v:"lt:18"`
	Gender string `json:"gender" v:"enum:female,male"`
}

func Example() {
	var s = Student{
		FName:  "John",
		LName:  "doe",
		Age:    18,
		Gender: "male",
	}

	v := validator.New()
	validation, err := v.ValidateStruct(s)
	if err != nil {
		panic(err)
	}

	fmt.Print(validation)
	// Output:
	// [{first_name [must be lowercase characters]} {Age [must be less than 18]}]
}

func NotEqual(val, param float64) error {
	if val == param {
		return fmt.Errorf("must be equal %v", param)
	}

	return nil
}

type Expire struct {
	Year int `json:"expire_year" v:"notequal:2017"`
}

func ExampleValidator_AddRuleNumeric() {
	var b = Expire{
		Year: 2017,
	}

	v := validator.New()
	err := v.AddRuleNumeric("notequal", NotEqual, 1)
	if err != nil {
		panic(err)
	}

	result, err := v.ValidateStruct(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// Output:
	// [{expire_year [must be equal 2017]}]
}

func StartCase(val string) error {
	pattern := `^([A-Z][a-z]*)+(\s[A-Z][a-z]*)*$`
	regex := regexp.MustCompile(pattern)

	if regex.MatchString(val) {
		return nil
	}

	return errors.New("must be start case")
}

type Book struct {
	Title string `json:"book_title" v:"startcases"`
}

func ExampleValidator_AddRuleString() {
	var b = Book{
		Title: "how To make Money",
	}

	v := validator.New()
	err := v.AddRuleString("startcases", StartCase)
	if err != nil {
		panic(err)
	}

	result, err := v.ValidateStruct(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// Output:
	// [{book_title [must be start case]}]
}
