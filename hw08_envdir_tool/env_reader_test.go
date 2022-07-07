package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadDir(t *testing.T) {
	t.Run("negative", func(t *testing.T) {
		env, err := ReadDir("fake/path")
		require.Error(t, err)
		require.Nil(t, env)

		env, err = ReadDir("testdata/env/BAR")
		require.Error(t, err)
		require.Nil(t, env)
	})

	t.Run("positive", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		expected := Environment{
			"BAR":   {Value: "bar", NeedRemove: false},
			"EMPTY": {Value: "", NeedRemove: false},
			"FOO":   {Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": {Value: "\"hello\"", NeedRemove: false},
			"UNSET": {Value: "", NeedRemove: true},
		}
		require.NoError(t, err)
		require.Equal(t, expected, env)
	})
}
