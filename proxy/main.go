package main

import (
	"io"
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			client, err := net.Dial("unix", "/tmp/mds.sock")
			if err != nil {
				panic(err)
			}
			defer client.Close()

			if _, err := io.Copy(conn, client); err != nil {
				log.Fatal(err)
			}
		}(conn)
	}

}
