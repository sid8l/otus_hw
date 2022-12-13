package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := map[string]struct {
		offset       int64
		limit        int64
		sourceString string
		result       []byte
	}{
		"empty file": {
			offset:       0,
			limit:        0,
			sourceString: "",
			result:       []byte(""),
		},
		"without limit and offset": {
			offset:       0,
			limit:        0,
			sourceString: "Hello World!",
			result:       []byte("Hello World!"),
		},
		"with offset": {
			offset:       6,
			limit:        0,
			sourceString: "Hello World!",
			result:       []byte("World!"),
		},
		"with limit": {
			offset:       0,
			limit:        5,
			sourceString: "Hello World!",
			result:       []byte("Hello"),
		},
		"with offset and limit": {
			offset:       3,
			limit:        5,
			sourceString: "Hello World!",
			result:       []byte("lo Wo"),
		},
		"limit > file size": {
			offset:       0,
			limit:        1000,
			sourceString: "Hello World!",
			result:       []byte("Hello World!"),
		},
		"with offset and limit > file size": {
			offset:       6,
			limit:        1000,
			sourceString: "Hello World!",
			result:       []byte("World!"),
		},
	}

	src, _ := os.CreateTemp("", "src")
	dst, _ := os.CreateTemp("", "dst")
	defer func() {
		os.Remove(src.Name())
		os.Remove(src.Name())
		src.Close()
		dst.Close()
	}()

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			src.Truncate(0)
			src.Seek(0, io.SeekStart)
			src.WriteString(test.sourceString)
			err := Copy(src.Name(), dst.Name(), test.offset, test.limit)
			require.Nil(t, err)
			dst.Seek(0, io.SeekStart)
			dstText, _ := io.ReadAll(dst)
			require.Equal(t, test.result, dstText)
		})
	}

	t.Run("copy from /dev/urandom", func(t *testing.T) {
		err := Copy("/dev/urandom", dst.Name(), 0, 500)
		require.Nil(t, err)
		dstStat, _ := dst.Stat()
		require.Equal(t, int64(500), dstStat.Size())
	})

	t.Run("negative tests", func(t *testing.T) {
		err := Copy(src.Name(), dst.Name(), 50, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)

		fakePath, _ := filepath.Abs("abracadabra")
		err = Copy(fakePath, dst.Name(), 0, 0)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})
}
