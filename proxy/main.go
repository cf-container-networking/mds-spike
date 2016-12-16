package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("mds-proxy <socket-path>")
		os.Exit(1)
	}

	socket := os.Args[1]

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
			client, err := net.Dial("unix", socket)
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
