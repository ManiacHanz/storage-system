package p2p

// Peer is an interfact that represents the remote node
type Peer interface{}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// form (TCP, UDP, websocket)
type Transport interface {
	ListenAndAccept() error
}
