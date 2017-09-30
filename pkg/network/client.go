package network

import uuid "github.com/satori/go.uuid"
import "net"

type Client struct {
	Username string
	UUID     uuid.UUID
	address  net.UDPAddr
}

// Create a new Client in challenge state
func NewClient(uuid uuid.UUID) *Client {

	return &Client{UUID: uuid}
}
