package network

// Netstack represents the servers networking stack
type NetStack struct {
	clients []Client
	port    int
}
