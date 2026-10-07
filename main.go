package main

import (
	"fmt"
)

func main() {
	v, err := OpenVolume("volume.img", 1024*1024*1024) // 1 GiB
	if err != nil {
		panic(err)
	}
	defer v.Close()

	data := make([]byte, 128*1024*1024)

	for i := range data {
		data[i] = 'A'
	}

	if err := v.WriteAt(0, data); err != nil {
		panic(err)
	}

	if err := v.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("write returned")

	result, err := v.ReadAt(4096, 5)
	if err != nil {
		panic(err)
	}

	fmt.Println("read:", string(result))
}
