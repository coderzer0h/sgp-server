package network

import uuid "github.com/satori/go.uuid"
import "net"

// Client represents a game client
type Client struct {
	Username string
	UUID     uuid.UUID
	address  net.UDPAddr
}

// NewClient returns a new Client
func NewClient(uuid uuid.UUID) *Client {

	return &Client{UUID: uuid}
}
