package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("mds-server <socket-path>")
		os.Exit(1)
	}

	socket := os.Args[1]

	l, err := net.ListenUnix("unix", &net.UnixAddr{socket, "unix"})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	defer os.Remove(socket)

	for {
		conn, err := l.AcceptUnix()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			_, err = conn.Write([]byte(`hello from mds\n`))
			if err != nil {
				panic(err)
			}
		}(conn)
	}

}
