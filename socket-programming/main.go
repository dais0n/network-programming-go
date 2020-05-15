package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/dais0n/network-programming-go/socket-programming/tcpclient"
	"github.com/dais0n/network-programming-go/socket-programming/tcpserver"
	"github.com/dais0n/network-programming-go/socket-programming/udpclient"
	"github.com/dais0n/network-programming-go/socket-programming/udpserver"
)

func main() {
	// get input value
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		os.Exit(1)
	}
	protocol := args[0]
	role := args[1]
	address := args[2]

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	// exec tcp or udp client, server
	switch protocol {
	case "tcp":
		switch role {
		case "server":
			err := tcpserver.Serve(address)
			if err != nil {
				log.Fatal(err.Error())
			}
		case "client":
			err := tcpclient.Connect(address)
			if err != nil {
				log.Fatal(err.Error())
			}
		default:
			missig_role()
		}
	case "udp":
		switch role {
		case "server":
			udpserver.Serve(address)
		case "client":
			udpclient.Connect(address)
		default:
			missig_role()
		}
	default:
		log.Fatal("Please specify tcp or udp on the 1st argument.")
	}
}

// missing role error
func missig_role() {
	log.Fatal("Please specify server or client on the 2nd argument.")
}
