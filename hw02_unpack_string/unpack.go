package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sourceString string) (string, error) {
	var resultString strings.Builder
	var previousLetter rune
	isPreviousLetter := false
	for i, r := range sourceString {
		if unicode.IsDigit(r) {
			if isPreviousLetter {
				multiplier, _ := strconv.Atoi(string(r))
				resultString.WriteString(strings.Repeat(string(previousLetter), multiplier))
				isPreviousLetter = false
			} else {
				return "", ErrInvalidString
			}
		} else {
			if isPreviousLetter {
				resultString.WriteRune(previousLetter)
			} else {
				isPreviousLetter = true
			}
			previousLetter = r
			if i == len(sourceString)-1 {
				resultString.WriteRune(r)
			}
		}
	}
	return resultString.String(), nil
}
