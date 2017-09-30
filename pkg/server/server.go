package server

import "github.com/coderzer0h/sgp-server/pkg/network"

// Server represents a SGP server
type Server struct {
	netstack *network.NetStack
	quit     chan bool
}

// NewServer generates and returns a new SGP server
func NewServer() *Server {
	n := network.NewNetStack()
	q := make(chan bool)

	s := Server{
		netstack: n,
		quit:     q,
	}
	return &s
}

// Start runs the SGP server
func (s *Server) Start() error {
	s.netstack.Listen()
	<-s.quit
	return nil
}
