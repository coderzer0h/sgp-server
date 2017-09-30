package network

import uuid "github.com/satori/go.uuid"

const (
	// Challenge represents the initial connection state
	Challenge = iota
	// Connecting represents the state in which a client begins to connect
	Connecting = iota
	// Connected represents the state in which a client is connected
	Connected = iota
)

// Session represents the connection between one client and the server
type Session struct {
	client *Client
	state  int
}

// NewSession creates a new session with a new client
func NewSession() *Session {

	client := NewClient(uuid.NewV4())
	return &Session{
		client: client,
		state:  Challenge,
	}
}
