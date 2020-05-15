package tcpclient

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func Connect(address string) error {
	// get input value
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		os.Exit(1)
	}
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	for {
		var sc = bufio.NewScanner(os.Stdin)
		var s string
		if sc.Scan() {
			s = sc.Text()
		}
		n, err := conn.Write([]byte(s))
		if err != nil {
			break
		}
		fmt.Printf("write byte length: %d\n", n)
		results := make([]byte, 1024)
		_, err = conn.Read(results)
		if err != nil {
			return err
		}

		fmt.Printf("read byte length: %d\n", n)
		fmt.Println(string(results))
	}
	return nil
}
