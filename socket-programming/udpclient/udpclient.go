package udpclient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Connect(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Println(err)
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
