package p2p

// Peer is a generic interface for a representation of a remote node, which is anything that can send and receive data
type Peer interface {

}

// Transport is a generic interface for a transport layer, which is anything that handles the communication between nodes in the network
// This can be iof the form of TCP, UDP, websockets, etc.
type Transport interface {
	ListenAndAccept() error

}
