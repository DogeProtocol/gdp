package interfaces

type OnReceiveMessage func([]byte, Node) error

type Gossip interface {
	Start(initialNodeList []Node, port uint16, keyPair SigKeyPair) error
	GossipMessage(message []byte) error
	AddNode([]Node) error
	OnReceiveMessage
	Stop() error
}
