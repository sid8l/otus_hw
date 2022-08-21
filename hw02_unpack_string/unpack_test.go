package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "a1b2x3", expected: "abbxxx"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "р1ус0ский2язык3", expected: "рускиййязыккк"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `\33npm`, expected: `333npm`},
		{input: `\3npm`, expected: `3npm`},
		{input: `\\3npm\\`, expected: `\\\npm\`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc",
		"45",
		"aaa10b",
		"a3b4c23",
		`d\n5abc`,
		`\n3abc`,
		`a3c2v4\t2`,
		`\fabc`,
		`asd\b`,
		`ab\amsd`,
		`\2\sasd\w`,
		`\qasdf\m\2`,
	}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
