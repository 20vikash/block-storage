package main

import (
	"fmt"
	"os"
)

const volumeSize = 1024 * 1024 * 1024 // 1 GiB

func flush(f *os.File) error {
	return f.Sync()
}

func writeAt(f *os.File, offset int64, data []byte) error {
	if offset < 0 || offset+int64(len(data)) > volumeSize {
		return fmt.Errorf("writeAt: offset out of bounds")
	}

	_, err := f.WriteAt(data, offset)
	return err
}

func readAt(f *os.File, offset int64, length int) ([]byte, error) {
	if offset < 0 || offset+int64(length) > volumeSize {
		return nil, fmt.Errorf("readAt: offset out of bounds")
	}

	buf := make([]byte, length)

	_, err := f.ReadAt(buf, offset)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func main() {
	f, err := os.OpenFile(
		"volume.img",
		os.O_CREATE|os.O_RDWR,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := f.Truncate(volumeSize); err != nil {
		panic(err)
	}

	data := make([]byte, 128*1024*1024)

	for i := range data {
		data[i] = 'A'
	}

	if err := writeAt(f, 0, data); err != nil {
		panic(err)
	}

	flush(f)

	fmt.Println("write returned")

	result, err := readAt(f, 4096, 5)
	if err != nil {
		panic(err)
	}

	fmt.Println("read:", string(result))
}
