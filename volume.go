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

func (v *Volume) WriteAt(offset int64, data []byte) error {
	if offset < 0 || offset+int64(len(data)) > v.size {
		return fmt.Errorf("writeAt: offset out of bounds")
	}

	_, err := v.file.WriteAt(data, offset)
	return err
}

func (v *Volume) ReadAt(offset int64, length int) ([]byte, error) {
	if offset < 0 || offset+int64(length) > v.size {
		return nil, fmt.Errorf("readAt: offset out of bounds")
	}

	buf := make([]byte, length)

	_, err := v.file.ReadAt(buf, offset)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (v *Volume) Flush() error {
	return v.file.Sync()
}

func (v *Volume) Close() error {
	return v.file.Close()
}
