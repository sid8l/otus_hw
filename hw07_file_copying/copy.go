package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	dst, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	srcStats, err := src.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	srcSize := srcStats.Size()

	if offset > srcSize {
		return ErrOffsetExceedsFileSize
	}

	srcReader := io.Reader(src)

	if limit == 0 {
		limit = srcSize
	} else {
		srcReader = io.LimitReader(srcReader, limit)
	}

	barLimit := limit
	if limit > srcSize-offset {
		barLimit = srcSize - offset
	}

	bar := pb.Full.Start64(barLimit)
	barReader := bar.NewProxyReader(srcReader)
	if _, err := src.Seek(offset, io.SeekStart); err != nil {
		return ErrUnsupportedFile
	}

	if _, err := io.Copy(dst, barReader); err != nil {
		return ErrOffsetExceedsFileSize
	}
	bar.Finish()

	return nil
}
