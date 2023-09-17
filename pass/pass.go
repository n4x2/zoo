// Package pass is password generator.
package pass

import (
	"crypto/rand"
	"errors"
	"math/big"
	"unicode"

	"github.com/n4x2/zoo/is"
)

const (
	defaultLength = 20
	minLength     = 1
	maxLength     = 50
)

var (
	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers = []byte("0123456789")
	symbols = []byte("!@#$%^&*")
)

// invalidLengthError error returns for invalid length password.
var invalidLengthError = errors.New("length must be between 1 and 50 characters")

// isConsecutiveType checks if 'p' and 'n' is same character type.
func isConsecutiveType(p, n rune) bool {
	return (unicode.IsLower(p) && unicode.IsLower(n)) ||
		(unicode.IsUpper(p) && unicode.IsUpper(n)) ||
		(unicode.IsNumber(p) && unicode.IsNumber(n)) ||
		(unicode.IsSymbol(p) && unicode.IsSymbol(n))
}

// randByte selects a random byte from a given slice of bytes.
func randByte(s []byte) byte {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(s))))
	if err != nil {
		panic(err)
	}
	return s[i.Int64()]
}

// Generate generates a random string.
//
// It returns a default 20-character random string containing letters, symbols and numbers.
// Symbols and numbers is optional, set 'n' to false to exclude numbers or 's' to false to
// exclude symbols.
//
// It will returns an error if provided length 'l' less than 1 or greater than 50.
func Generate(n, s bool, l ...int) (string, error) {
	var length = defaultLength
	if len(l) > 0 {
		length = l[0]
	}

	if length < minLength || length > maxLength {
		return "", invalidLengthError
	}

	var chars = letters
	if n {
		chars = append(chars, numbers...)
	}

	if s {
		chars = append(chars, symbols...)
	}

	var p = make([]byte, length)
	for i := range p {
		for {
			var c = randByte(chars)
			if i == 0 || (!isConsecutiveType(rune(p[i-1]), rune(c)) && !is.Contain(p[:i], c)) {
				p[i] = c
				break
			}

			if s {
				if !is.ContainOneOf(p[:i], symbols) {
					p[i] = randByte(symbols)
					s = false
					break
				}
			}

			if n {
				if !is.ContainOneOf(p[:i], numbers) {
					p[i] = randByte(numbers)
					n = false
					break
				}
			}
		}
	}
	return string(p), nil
}
