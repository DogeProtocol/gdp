package interfaces

type OnNodeDiscover func(Node) error

type Discovery interface {
	Start(seedNodes []Node, port uint64, keyPair SigKeyPair) error
	OnNodeDiscover
	Stop() error
}
