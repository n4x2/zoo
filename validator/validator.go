// Package validator provides struct validation based on tags.
package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/n4x2/zoo/is"
	"github.com/n4x2/zoo/regex"
	"github.com/n4x2/zoo/to"
)

// Default values for validation.
const (
	ValidatorTag = "v"    // Default validator tag.
	SkipTag      = "-"    // Default tag to skip validation.
	JSONTag      = "json" // Default JSON tag.
	PairSep      = ":"    // Default pair separator.
	ParamSep     = ","    // Default parameter separator.
	TagSep       = "|"    // Default tag separator.
)

// Index for tag parsing.
const (
	NameIndex  = 0 // Index for tag name.
	ParamIndex = 1 // Index for tag parameter.
)

// E store associates error messages with specific validation rules.
var E = map[string]string{
	"alpha":     "must be alphabetic characters",
	"alphadash": "must be alphaNeric characters, dash, and underscore",
	"alphanum":  "must be alphanumeric characters",
	"ascii":     "must be ASCII characters",
	"enum":      "%v not allowed for this field",
	"email":     "invalid email address",
	"equal":     "must be the same as %v",
	"gt":        "must be greater than %v",
	"gte":       "must be greater than or equal to %v",
	"lat":       "invalid latitude: %v",
	"lon":       "invalid longitude: %v",
	"lt":        "must be less than %v",
	"lte":       "must be less than or equal to %v",
	"lowercase": "must be lowercase characters",
	"range":     "value must be in range %v-%v",
	"ulid":      "invalid ULID",
	"uuid":      "invalid UUID",
	"uppercase": "must be uppercase characters",
}

// R stores default validation tags, it wraps functions from the [is]
// package to perform validation.
//
// [is]: https://pkg.go.dev/github.com/n4x2/zoo/is
var R = map[string]Detail{
	"alpha":     {Fn: is.Alpha, Maxp: 0, N: false},
	"alphadash": {Fn: is.AlphaDash, Maxp: 0, N: false},
	"alphanum":  {Fn: is.AlphaNumeric, Maxp: 1, N: false},
	"ascii":     {Fn: is.ASCII, Maxp: 0, N: false},
	"enum":      {Fn: is.Contain[[]string, string], Maxp: -1, N: false},
	"email":     {Fn: is.Email, Maxp: 0, N: false},
	"equal":     {Fn: is.Equal[float64], Maxp: 1, N: true},
	"gt":        {Fn: is.GreaterThan[float64], Maxp: 1, N: true},
	"gte":       {Fn: is.GreaterThanEqual[float64], Maxp: 1, N: true},
	"lat":       {Fn: is.Latitude, Maxp: 0, N: false},
	"lon":       {Fn: is.Longitude, Maxp: 0, N: false},
	"lt":        {Fn: is.LessThan[float64], Maxp: 1, N: true},
	"lte":       {Fn: is.LessThanEqual[float64], Maxp: 1, N: true},
	"lowercase": {Fn: is.Lowercase, Maxp: 0, N: false},
	"range":     {Fn: is.Range[float64], Maxp: 2, N: true},
	"ulid":      {Fn: is.ULID, Maxp: 0, N: false},
	"uuid":      {Fn: is.UUID, Maxp: 0, N: false},
	"uppercase": {Fn: is.Uppercase, Maxp: 0, N: false},
}

// Error variables for common error conditions that may be
// encountered during validation.
var (
	// Indicates that an unexported field was encountered during
	// validation.
	errUnexportedField = errors.New("unexported field encountered")
	// Indicates that the input is not a struct.
	errInvalidInput = errors.New("input is not a struct")
	// Indicates that validator tag is missing on a struct field.
	errMissingTag = errors.New("missing validator tag")
	// Indicates that a validator tag is not supported.
	errTagUnsupported = errors.New("tag is not supported")
	// Indicates that a parameter is not allowed for a specific
	// validator tag.
	errParamNotAllowed = errors.New("parameter not allowed")
)

type (
	// Detail holds the tag validation details.
	Detail struct {
		Fn   any  // Function for validation.
		Maxp int  // Maximum allowed parameter.
		N    bool // Set 'true' if tag processing numerical value.
	}

	// Field represents fields data containing name, value, and
	// associated validation tags.
	Field struct {
		N string // The field name.
		V any    // The field value.
		T []Tag  // Validation tags.
	}

	// Tag represents a validation tag, including tag name and
	// optional parameters.
	Tag struct {
		N string // The tag name.
		P []any  // Optional parameters.
	}
)

type (
	// Result represents a validation error, including the field
	// name and error messages.
	Result struct {
		F string   // The field name.
		E []string // Error messages.
	}

	// Validator contains default error messages and validation
	// rules.
	Validator struct {
		msg   map[string]string
		rules map[string]Detail
	}
)

type (
	// errInvalidParam an error type for cases where an invalid
	// parameter is encountered.
	errInvalidParam struct {
		tn string // The tag name.
		v  any    // The value that is not accepted.
	}

	// errTypeConversion an error type for cases where fail type
	// conversion is encountered.
	errTypeConversion struct {
		tn string // The tag name.
		t  string // Target type e.g "string".
		v  any    // The value.
	}
)

// Error an error for the errInvalidParam type.
func (e *errInvalidParam) Error() string {
	return fmt.Sprintf("tag %s only accept %v parameter", e.tn, e.v)
}

// Error an error for the errTypeConversion type.
func (e *errTypeConversion) Error() string {
	return fmt.Sprintf("%s: fail to convert %v to %s", e.tn, e.v, e.t)
}

// parseParam parse tag parameter value and convert it into one of this
// possible types: float64 and string, based on the value format. It
// returns an error if parsing fails.
func parseParam(v string) (any, error) {
	switch {
	case regex.Numeric.MatchString(v):
		return to.Float64(v)
	default:
		return v, nil
	}
}

// parseTag parse tag name and parameters. It returns an error if parsing
// parameter value fails
func parseTag(v string) ([]Tag, error) {
	var s = strings.Split(v, TagSep)

	var t = make([]Tag, len(s))
	for i, tval := range s {
		var tp = strings.Split(tval, PairSep)

		t[i].N = tp[NameIndex]

		m, ok := R[t[i].N]
		if !ok && t[i].N != SkipTag {
			return nil, fmt.Errorf("%w: %s", errTagUnsupported, t[i].N)
		}

		if m.Maxp == 0 && len(tp) > 1 {
			return nil, fmt.Errorf("%w: %s", errParamNotAllowed, t[i].N)
		}

		if len(tp) > 1 {
			var pv = strings.Split(tp[ParamIndex], ParamSep)
			if len(pv) != m.Maxp && m.Maxp != -1 {
				return nil, &errInvalidParam{tn: t[i].N, v: m.Maxp}
			}

			for _, pval := range pv {
				if m.N && !regex.Numeric.MatchString(pval) {
					return nil, &errInvalidParam{tn: t[i].N, v: "numeric"}
				}

				val, err := parseParam(pval)
				if err != nil {
					return nil, err
				}
				t[i].P = append(t[i].P, val)
			}
		}
	}
	return t, nil
}

// serialize parse value of struct fields. It return an error if input is
// not a struct, field is unexported, missing validator tag, or fail to
// parse tag.
func serialize(v any) ([]Field, error) {
	if !is.Struct(v) {
		return nil, errInvalidInput
	}

	var s = reflect.ValueOf(v)
	var t = s.Type()
	var n = s.NumField()

	var f = make([]Field, n)
	for i := 0; i < n; i++ {
		var fi = t.Field(i)

		f[i].N = fi.Name
		if fi.PkgPath != "" {
			return nil, fmt.Errorf("field %s : %w", f[i].N, errUnexportedField)
		}

		if jn := fi.Tag.Get(JSONTag); jn != "" {
			f[i].N = jn
		}

		ft, ok := fi.Tag.Lookup(ValidatorTag)
		if ok && ft == "" {
			return nil, fmt.Errorf("field %s: %w", f[i].N, errMissingTag)
		}

		pt, err := parseTag(ft)
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", f[i].N, err)
		}
		f[i].T = pt
		f[i].V = s.Field(i).Interface()
	}
	return f, nil
}

// AddRuleNumeric add custom rule that processing numeric values into validator.
// It need 'n' tag name, 'fn' function that perform validation, and 'maxp' to
// determine maximum allowed parameter. It returns an error if tag name already
// exists.
func (r *Validator) AddRuleNumeric(n string, fn func(float64, float64) error, maxp int) error {
	if _, exists := r.rules[n]; exists {
		return errors.New("rule" + n + "already exists")
	}

	r.rules[n] = Detail{
		Fn:   fn,
		Maxp: maxp,
		N:    true,
	}
	return nil
}

// AddRuleString add custom rule that processing string values into validator.
// It require 'n' tag name and 'fn' function that perform validation as input.
// It returns an error if tag name already exists.
func (r *Validator) AddRuleString(n string, fn func(string) error) error {
	if _, exists := r.rules[n]; exists {
		return errors.New("rule" + n + "already exists")
	}

	r.rules[n] = Detail{
		Fn:   fn,
		Maxp: 0,
		N:    false,
	}
	return nil
}

// ValidateField validate given value based on tags. It returns slices of
// validation messages if any validation error encountered. It returns an
// error if rule is not found or type conversion is failed.
func (r *Validator) ValidateField(v any, st []Tag) ([]string, error) {
	var e = make([]string, 0)

	for _, t := range st {
		if t.N == SkipTag && v == nil {
			continue
		}

		m, ok := r.rules[t.N]
		if !ok && t.N != SkipTag {
			return nil, fmt.Errorf("%s: %w", t.N, errTagUnsupported)
		}

		switch fn := m.Fn.(type) {
		case func(string) bool:
			val, ok := v.(string)
			if !ok {
				return nil, &errTypeConversion{tn: t.N, t: "string", v: v}
			}

			if !fn(val) {
				e = append(e, r.msg[t.N])
			}
		case func(float64, float64) bool:
			val, err := to.Float64(v)
			if err != nil {
				return nil, &errTypeConversion{tn: t.N, t: "float64", v: v}
			}

			for _, tp := range t.P {
				p, err := to.Float64(tp)
				if err != nil {
					return nil, &errTypeConversion{tn: t.N, t: "float64", v: v}
				}

				if !fn(val, p) {
					e = append(e, fmt.Sprintf(r.msg[t.N], p))
					break
				}
			}
		case func([]string, string) bool:
			var p = make([]string, len(t.P))
			for i, tp := range t.P {
				val, ok := tp.(string)
				if !ok {
					return nil, &errTypeConversion{tn: t.N, t: "string", v: v}
				}
				p[i] = val
			}

			val, ok := v.(string)
			if !ok {
				return nil, &errTypeConversion{tn: t.N, t: "string", v: v}
			}

			if !fn(p, val) {
				e = append(e, fmt.Sprintf(r.msg[t.N], v))
			}
		case func(string) error:
			err := fn(v.(string))
			if err != nil {
				e = append(e, err.Error())
			}
		case func(float64, float64) error:
			val, err := to.Float64(v)
			if err != nil {
				return nil, &errTypeConversion{tn: t.N, t: "float64", v: v}
			}

			for _, tp := range t.P {
				p, err := to.Float64(tp)
				if err != nil {
					return nil, &errTypeConversion{tn: t.N, t: "float64", v: v}
				}

				err = fn(val, p)
				if err != nil {
					e = append(e, err.Error())
					break
				}
			}
		}
	}
	return e, nil
}

// ValidateStruct validate given struct based on their associated tags.
// It will returns slices of [Result] containing field name and error
// messages if any validation error encountered. It returns an error if
// input is not struct or failed to parse fields or failed to convert
// values.
func (r *Validator) ValidateStruct(v any) ([]Result, error) {
	pf, err := serialize(v)
	if err != nil {
		return nil, fmt.Errorf("validator: %w", err)
	}

	var res = make([]Result, 0)
	for _, f := range pf {
		m, err := r.ValidateField(f.V, f.T)
		if err != nil {
			return nil, fmt.Errorf("validator: %w", err)
		}

		if len(m) > 0 {
			res = append(res, Result{F: f.N, E: m})
		}
	}
	return res, nil
}

// New creates new validator instances.
func New() *Validator {
	return &Validator{
		msg:   E,
		rules: R,
	}
}
