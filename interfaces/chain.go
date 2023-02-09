package interfaces

type Chain interface {
	ID() string
	RootHash() string
	LatestBlockNumber() uint64
	LatestBlockHash() string
}
