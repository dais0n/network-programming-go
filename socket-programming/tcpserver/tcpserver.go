package tcpserver

import (
	"fmt"
	"net"
)

func Serve(address string) error {
	// get listner socket
	ln, err := net.Listen("tcp", address)
	defer ln.Close()

	if err != nil {
		// handle error
		return err
	}
	for {
		// get connected socket
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil {
			// handle error
			return err
		}

		handler(conn)
	}
}

func handler(conn net.Conn) error {
	fmt.Printf("Handling data from %s\n", conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("Connection closed.")
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(string(buf[:n]))
		conn.Write(buf[:n])
	}
	return nil
}
