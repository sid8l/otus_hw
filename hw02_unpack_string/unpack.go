package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const shielding = '\\'

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	var resultString strings.Builder
	ra := []rune(inputString)
	isCharacterEscaped := false
	for i := range ra {
		if unicode.IsDigit(ra[i]) && (i == 0 || unicode.IsDigit(ra[i-1]) && ra[i-2] != shielding) {
			return "", ErrInvalidString
		} else if unicode.IsDigit(ra[i]) && !isCharacterEscaped {
			continue
		}
		if ra[i] == shielding {
			if i != len(ra)-1 && !unicode.IsDigit(ra[i+1]) && ra[i+1] != shielding {
				return "", ErrInvalidString
			} else if !isCharacterEscaped {
				isCharacterEscaped = true
				continue
			}
		}

		if i < len(ra)-1 && unicode.IsDigit(ra[i+1]) {
			nextNumber, err := strconv.Atoi(string(ra[i+1]))
			if err != nil {
				return "", err
			}
			resultString.WriteString(strings.Repeat(string(ra[i]), nextNumber))
		} else {
			resultString.WriteRune(ra[i])
		}
		isCharacterEscaped = false
	}
	if isCharacterEscaped {
		return "", ErrInvalidString
	}
	return resultString.String(), nil
}
