package network

import (
	"bytes"
	"net"
	"sync"

	"github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
	"github.com/vmihailenco/msgpack"
)

// NetStack represents the servers networking stack
type NetStack struct {
	address  net.UDPAddr
	sessions map[uuid.UUID]Session
	port     int
}

// NewNetStack returns a new networking stack
func NewNetStack() *NetStack {
	s := make(map[uuid.UUID]Session, 10)
	n := NetStack{
		sessions: s,
		port:     31337,
	}
	return &n
}

// Listen listens for connections
func (n *NetStack) Listen() error {
	addr, err := net.ResolveUDPAddr("udp", ":31337")
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		var wg sync.WaitGroup
		go func() {
			wg.Add(1)
			logrus.Infof("Listening on port %s", "31337")
			ln, err := net.ListenUDP("udp", addr)
			if err != nil {
				logrus.Fatal(err)
			}

			buf := make([]byte, 10240)

			for {
				n, addr, err := ln.ReadFromUDP(buf)
				if err != nil {
					logrus.Error(err)
				}

				r := bytes.NewReader(buf)
				mp := msgpack.NewDecoder(r)

				id, err := mp.DecodeInt()
				if err != nil {
					logrus.Info(err)
				}
				seq, err := mp.DecodeInt64()
				if err != nil {
					logrus.Info(err)
				}

				msg, err := mp.DecodeString()
				if err != nil {
					logrus.Info(err)
				}

				logrus.Infof("Read %s,%s,%s,%s from %s", id, seq, msg, n, addr)
			}

		}()
		wg.Wait()
	}()
	return nil
}
