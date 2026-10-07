package main

import (
	"fmt"
	"os"
)

type Volume struct {
	file *os.File
	size int64
}

func OpenVolume(path string, size int64) (*Volume, error) {
	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_RDWR,
		0644,
	)
	if err != nil {
		return nil, err
	}

	if err := f.Truncate(size); err != nil {
		f.Close()
		return nil, err
	}

	return &Volume{
		file: f,
		size: size,
	}, nil
}

func (v *Volume) WriteAt(p []byte, offset int64) (int, error) {
	if offset < 0 || offset+int64(len(p)) > v.size {
		return 0, fmt.Errorf("write out of bounds")
	}

	return v.file.WriteAt(p, offset)
}

func (v *Volume) ReadAt(p []byte, offset int64) (int, error) {
	if offset < 0 || offset+int64(len(p)) > v.size {
		return 0, fmt.Errorf("readAt: offset out of bounds")
	}

	return v.file.ReadAt(p, offset)
}

func (v *Volume) Flush() error {
	return v.file.Sync()
}

func (v *Volume) Close() error {
	return v.file.Close()
}
