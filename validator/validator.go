// Package validator provides struct value validation based on tags.
package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	indexName  = 0 // Index of tag name.
	indexParam = 1 // Index of parameter.
)

const (
	jsonTag           = "json"     // JSON tag.
	pairSeparatorTag  = ":"        // Separator for tag name and parameters.
	paramSeparatorTag = ","        // Separator for parameter.
	separatorTag      = "|"        // Separator for multiple tags.
	skipValidationTag = "-"        // Skip validation tag
	validationTag     = "validate" // Default validation tag.
)

const (
	alphaPattern = "^[a-zA-Z]+$" // Alphabet pattern.
)

// Regexes for validation.
var (
	AlphaRegex = regexp.MustCompile(alphaPattern) // Alphabet regex a-z or A-Z.
)

var (
	// errInvalidStruct error for non struct input.
	errInvalidStruct = errors.New("input must be a struct")
	// errUnexportedField error for unexported field.
	errUnexportedField = errors.New("unexported field encountered")
)

// Rule represents rule for validation.
type Rule interface {
	Name() string
	Validate(f, v, p string) error
}

type (
	enum     struct{}
	required struct{}

	// TODO
	// add more rules.
)

func (r *enum) Name() string     { return "enum" }
func (r *required) Name() string { return "required" }

func (r *enum) Validate(_, v, p string) error {
	options := strings.Split(p, paramSeparatorTag)
	// TODO
	// use slices.Contains() in go 1.21.
	for _, o := range options {
		if v != o {
			return fmt.Errorf("selected %s is invalid", v)
		}
	}

	return nil
}

func (r *required) Validate(f, v, _ string) error {
	if v == "" {
		return fmt.Errorf("%s is required", f)
	}
	return nil
}

type (
	// ErrValidation contains field name and validation errors.
	ErrValidation struct {
		Field string
		Err   []string
	}

	// Field contains name, value, and tags of the field.
	Field struct {
		Name, Value string
		Tags        []Tag
	}

	// Tag contain name tag and the parameters.
	Tag struct {
		Name  string
		Param string
	}

	// Validator contains validator tag and default rules.
	Validator struct {
		validatorTag string
		rules        []Rule
	}
)

// isStructType check if given value is a struct.
func isStructType(s interface{}) bool {
	v := reflect.ValueOf(s)

	return v.Kind() == reflect.Struct
}

// parseName parse field name.
func parseName(f reflect.StructField) string {
	jsonName, exist := f.Tag.Lookup(jsonTag)
	if exist {
		return jsonName
	}

	return f.Name
}

// parseValue parse field value.
func parseValue(v reflect.Value) string {
	if v.IsZero() {
		return ""
	}

	return fmt.Sprintf("%v", v.Interface())
}

// parseTag parse field tags.
func parseTag(v string) []Tag {
	var tags []Tag

	listOfTag := strings.Split(v, separatorTag)
	for _, t := range listOfTag {
		var tag Tag

		partsOfTag := strings.Split(t, pairSeparatorTag)
		if len(partsOfTag) > 1 {
			tag.Param = partsOfTag[indexParam]
		}
		tag.Name = partsOfTag[indexName]

		tags = append(tags, tag)
	}

	return tags
}

// AddRule add custom rule into validator, it will return
// errors if the rule (tag) already exists.
func (v *Validator) AddRule(r Rule) error {
	for _, rule := range v.rules {
		if r.Name() == rule.Name() {
			return fmt.Errorf("validator: rule %s already exist", r.Name())
		}
	}
	v.rules = append(v.rules, r)

	return nil
}

// serialize serialize struct into fields, it will return
// error if unexported field encountered.
func (v *Validator) serialize(s interface{}) ([]Field, error) {
	var fields []Field

	val := reflect.ValueOf(s)
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		// Checks field is exportable.
		if t.Field(i).PkgPath != "" {
			return nil, errUnexportedField
		}

		var f Field

		f.Name = parseName(t.Field(i))
		f.Value = parseValue(val.Field(i))
		f.Tags = parseTag(t.Field(i).Tag.Get(v.validatorTag))

		fields = append(fields, f)
	}

	return fields, nil
}

// SetValidatorTag set custom validator tag, it will return
// error if given value is not alphabetic characters.
func (v *Validator) SetValidatorTag(n string) error {
	if !AlphaRegex.MatchString(n) {
		return fmt.Errorf("validator: tag %s must be letters", n)
	}
	v.validatorTag = n

	return nil
}

// validateField validate the field.
func (v *Validator) validateField(f Field) ErrValidation {
	var e ErrValidation
	for _, t := range f.Tags {
		if t.Name == skipValidationTag && f.Value == "" {
			break
		}

		for _, r := range v.rules {
			if r.Name() == t.Name {
				err := r.Validate(f.Name, f.Value, t.Param)
				if err != nil {
					e.Field = f.Name
					e.Err = append(e.Err, err.Error())
				}
			}
		}
	}

	return e
}

// Validate validate given struct based on associated tags.
func (v *Validator) Validate(s interface{}) ([]ErrValidation, error) {
	if ok := isStructType(s); !ok {
		return nil, fmt.Errorf("validator: %s", errInvalidStruct)
	}

	fields, err := v.serialize(s)
	if err != nil {
		return nil, fmt.Errorf("validator: %s", err)
	}

	var errs []ErrValidation

	for _, f := range fields {
		err := v.validateField(f)
		if err.Err != nil {
			errs = append(errs, err)
		}
	}

	return errs, nil
}

// New create new validator instances.
func New() *Validator {
	defaultRule := []Rule{
		&enum{},
		&required{},
	}

	return &Validator{
		validatorTag: validationTag,
		rules:        defaultRule,
	}
}
