package packets

import "github.com/google/gopacket/layers"

type GettableEndPoints interface {
	GetSource() string
	GetDistination() string
	GetPayload() []byte
}

type GettableEndPointsTCP struct {
	TCP *layers.TCP
}

func (g *GettableEndPointsTCP) GetSource() string {
	return g.TCP.SrcPort.String()
}

func (g *GettableEndPointsTCP) GetDistination() string {
	return g.TCP.DstPort.String()
}

func (g *GettableEndPointsTCP) GetPayload() []byte {
	return g.TCP.Payload
}

type GettableEndPointsUDP struct {
	UDP *layers.UDP
}

func (g *GettableEndPointsUDP) GetSource() string {
	return g.UDP.SrcPort.String()
}

func (g *GettableEndPointsUDP) GetDistination() string {
	return g.UDP.DstPort.String()
}

func (g *GettableEndPointsUDP) GetPayload() []byte {
	return g.UDP.Payload
}
