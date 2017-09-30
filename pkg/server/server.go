package server

import "github.com/coderzer0h/sgp-server/pkg/network"

// Server represents a SGP server
type Server struct {
	netstack network.NetStack
}

// NewServer generates and returns a new SGP server
func NewServer() (server *Server) {

	return
}

// Start runs the SGP server
func (s *Server) Start() error {
	return nil
}
