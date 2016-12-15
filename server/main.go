package main

import (
	"net"
	"os"
)

func main() {
	l, err := net.ListenUnix("unix", &net.UnixAddr{"/tmp/mds.sock", "unix"})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	defer os.Remove("/tmp/mds.sock")

	for {
		conn, err := l.AcceptUnix()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			b := []byte(`hello`)
			_, err = conn.Write(b)
			if err != nil {
				panic(err)
			}
		}(conn)
	}

}
