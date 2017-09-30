package network

import (
	"net"
	"sync"

	"github.com/Sirupsen/logrus"
)

// NetStack represents the servers networking stack
type NetStack struct {
	clients []Client
	port    int
}

// NewNetStack returns a new networking stack
func NewNetStack() *NetStack {
	c := make([]Client, 10)
	n := NetStack{
		clients: c,
		port:    1337,
	}
	return &n
}

// Listen listens for connections
func (n *NetStack) Listen() error {
	addr, err := net.ResolveUDPAddr("udp", ":11337")
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		var wg sync.WaitGroup
		go func() {
			wg.Add(1)
			logrus.Infof("Listening on port %s", "11337")
			ln, err := net.ListenUDP("udp", addr)
			if err != nil {
				logrus.Fatal(err)
			}

			buf := make([]byte, 1024)

			for {
				n, addr, err := ln.ReadFromUDP(buf)
				if err != nil {
					logrus.Error(err)
				}
				logrus.Infof("Read %s from %s", n, addr)
			}

		}()
		wg.Wait()
	}()
	return nil
}
