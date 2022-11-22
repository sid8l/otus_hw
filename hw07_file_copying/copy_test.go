package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	src, _ := os.CreateTemp("", "src")
	dst, _ := os.CreateTemp("", "dst")
	defer func() {
		os.Remove(src.Name())
		os.Remove(src.Name())
		src.Close()
		dst.Close()
	}()

	t.Run("positive tests", func(t *testing.T) {
		err := Copy(src.Name(), dst.Name(), 0, 0)
		require.Nil(t, err)
		dstText, _ := io.ReadAll(dst)
		require.Equal(t, []byte(""), dstText)

		src.WriteString("Hello World!")

		err = Copy(src.Name(), dst.Name(), 0, 0)
		require.Nil(t, err)
		dstText, _ = io.ReadAll(dst)
		require.Equal(t, []byte("Hello World!"), dstText)

		err = Copy(src.Name(), dst.Name(), 6, 0)
		require.Nil(t, err)
		dst.Seek(0, io.SeekStart)
		dstText, _ = io.ReadAll(dst)
		require.Equal(t, []byte("World!"), dstText)

		err = Copy(src.Name(), dst.Name(), 0, 5)
		require.Nil(t, err)
		dst.Seek(0, io.SeekStart)
		dstText, _ = io.ReadAll(dst)
		require.Equal(t, []byte("Hello"), dstText)

		err = Copy(src.Name(), dst.Name(), 3, 5)
		require.Nil(t, err)
		dst.Seek(0, io.SeekStart)
		dstText, _ = io.ReadAll(dst)
		require.Equal(t, []byte("lo Wo"), dstText)

		err = Copy(src.Name(), dst.Name(), 0, 1000)
		require.Nil(t, err)
		dst.Seek(0, io.SeekStart)
		dstText, _ = io.ReadAll(dst)
		require.Equal(t, []byte("Hello World!"), dstText)

		err = Copy("/dev/urandom", dst.Name(), 0, 500)
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
