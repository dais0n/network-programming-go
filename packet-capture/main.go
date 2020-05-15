package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const WIDTH = 20

func main() {
	// get input value
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("")
	}
	inputDevice := args[0]

	// get the input network interface
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	var flag bool
	for _, device := range devices {
		if inputDevice == device.Name {
			flag = true
			break
		}
	}

	if !flag {
		log.Fatal("Failed to get interface")
	}

	// open device
	handle, err := pcap.OpenLive(
		inputDevice,
		1024,
		false,
		30*time.Second,
	)
	if err != nil {
		log.Fatal(err)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		frame, _ := ethernetLayer.(*layers.Ethernet)
		switch frame.NextLayerType() {
		case layers.LayerTypeIPv4:
			ipv4Handler(frame)
		}
	}
}

func printPacketInfo(l3 *layers.IPv4, l4 *layers.TCP, proto string) {
	fmt.Printf(
		"Captured a %s packet from %s|%s to %s|%s\n",
		proto,
		l3.SrcIP,
		l4.SrcPort,
		l3.DstIP,
		l4.DstPort,
	)

	payload := l4.Payload

	for i := 0; i < len(payload); i++ {
		fmt.Printf("%02X ", payload[i])
		if i%WIDTH == WIDTH-1 || i == len(payload)-1 {
			for j := 0; j < WIDTH-1-(i%(WIDTH)); j++ {
				fmt.Print("   ")
			}
			fmt.Print("| ")
			for k := i - i%WIDTH; k < i+1; k++ {
				if isASCIIAlphabetic(payload[k]) {
					fmt.Print(string(payload[k]))
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
	fmt.Printf(strings.Repeat("=", WIDTH*3))
	fmt.Println()
}

func ipv4Handler(ethernetPacket *layers.Ethernet) {
	packet := gopacket.NewPacket(ethernetPacket.Payload, layers.LayerTypeIPv4, gopacket.Default)
	ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
	ipv4Packet, _ := ipv4Layer.(*layers.IPv4)
	switch ipv4Packet.NextLayerType() {
	case layers.LayerTypeTCP:
		tcp_handler(ipv4Packet)
	case layers.LayerTypeUDP:
		udp_handler(ipv4Packet)
	default:
		fmt.Println("Not a tcp or udp packet")
	}
}

func tcp_handler(ipv4Packet *layers.IPv4) {
	packet := gopacket.NewPacket(ipv4Packet.Payload, layers.LayerTypeTCP, gopacket.Default)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	tcpPacket, _ := tcpLayer.(*layers.TCP)
	printPacketInfo(ipv4Packet, tcpPacket, "TCP")
}

func udp_handler(ipv4Packet *layers.IPv4) {
	// packet := gopacket.NewPacket(ipv4Packet.Payload, layers.LayerTypeUDP, gopacket.Default)
	// udpLayer := packet.Layer(layers.LayerTypeUDP)
	// udpPacket, _ := udpLayer.(*layers.UDP)
	//printPacketInfo(ipv4Packet, udpPacket, "TCP")
}

func isASCIIAlphabetic(s byte) bool {
	if 65 <= s && s <= 90 {
		return true
	}
	if 97 <= s && s <= 122 {
		return true
	}
	return false
}
