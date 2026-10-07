package main

import (
	"fmt"
	"net"

	"github.com/pojntfx/go-nbd/pkg/server"
)

func main() {
	volume, err := OpenVolume("volume.img", 1024*1024*1024)
	if err != nil {
		panic(err)
	}
	defer volume.Close()

	backend := NewNBDBackend(volume)

	listener, err := net.Listen("tcp", "127.0.0.1:10809")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("NBD server listening on 127.0.0.1:10809")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()

			err := server.Handle(
				conn,
				[]*server.Export{
					{
						Name:        "volume",
						Description: "Our first block store",
						Backend:     backend,
					},
				},
				&server.Options{
					MinimumBlockSize:   512,
					PreferredBlockSize: 4096,
					MaximumBlockSize:   4096,
				},
			)
			if err != nil {
				fmt.Println("NBD connection error:", err)
			}
		}()
	}
}
