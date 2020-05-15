package udpserver

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Serve(address string) error {
	addressInfos := strings.Split(address, ":")

	ip := net.ParseIP(addressInfos[0])
	port, err := strconv.Atoi(addressInfos[1])

	if err != nil {
		return err
	}

	// one thread
	udpLn, err := net.ListenUDP(
		"udp",
		&net.UDPAddr{
			IP:   ip,
			Port: port,
		},
	)
	if err != nil {
		return err
	}
	defer udpLn.Close()

	buf := make([]byte, 1024)
	for {
		_, addr, err := udpLn.ReadFromUDP(buf)
		if err != nil {
			return err
		}
		fmt.Println(addr)
		fmt.Println(string(buf))
		udpLn.WriteTo(buf, addr)
	}
}
