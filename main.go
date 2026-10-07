package main

import "fmt"

func main() {
	v, err := OpenVolume("volume.img", 1024*1024*1024)
	if err != nil {
		panic(err)
	}
	defer v.Close()

	data := []byte("HELLO")

	_, err = v.WriteAt(data, 4096)
	if err != nil {
		panic(err)
	}

	if err := v.Flush(); err != nil {
		panic(err)
	}

	result := make([]byte, 5)

	_, err = v.ReadAt(result, 4096)
	if err != nil {
		panic(err)
	}

	fmt.Println("read:", string(result))
}
