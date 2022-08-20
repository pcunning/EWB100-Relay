package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"github.com/dmichael/go-multicast/multicast"
	"github.com/nats-io/nats.go"
)

const (
	maxDatagramSize = 8192
	addr            = "239.192.2.2:5000"
)

type Router struct {
	nc *nats.Conn
	mc *net.UDPConn
}

func (r *Router) handleUDP(src *net.UDPAddr, n int, b []byte) {
	log.Printf("UDP: %s %s\n", src, hex.Dump(b[:n]))
	r.nc.Publish("ch1", b[:n])
}

func (r *Router) handleNATS(m *nats.Msg) {
	fmt.Printf("NATS: %s\n", string(m.Data))
	r.mc.Write(m.Data)
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	router := &Router{
		nc: nc,
		mc: conn,
	}

	nc.Subscribe("ch1", router.handleNATS)
	multicast.Listen(addr, router.handleUDP)
}
