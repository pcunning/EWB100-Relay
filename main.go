package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"github.com/dmichael/go-multicast/multicast"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

const (
	maxDatagramSize = 8192
	addr            = "239.192.2.2:5000"
)

type Router struct {
	uuid  string
	local net.Addr
	nc    *nats.Conn
	mc    *net.UDPConn
}

func (r *Router) handleUDP(src *net.UDPAddr, n int, b []byte) {
	if src != r.local {
		log.Printf("UDP: %s %s\n", src, hex.Dump(b[:n]))
		r.nc.PublishMsg(&nats.Msg{
			Subject: "ch1",
			Data:    b[:n],
			Header: map[string][]string{
				"uuid": []string{r.uuid},
			},
		})
	}
}

func (r *Router) handleNATS(m *nats.Msg) {
	fmt.Printf("NATS: %s\n", string(m.Data))
	if m.Header["uuid"][0] != r.uuid {
		r.mc.Write(m.Data)
	}
}

func main() {
	nc, _ := nats.Connect("nats://demo.nats.io:4222")
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	router := &Router{
		uuid:  uuid.New().String(),
		local: conn.LocalAddr(),
		nc:    nc,
		mc:    conn,
	}

	nc.Subscribe("ch1", router.handleNATS)
	multicast.Listen(addr, router.handleUDP)
}
