package interfaces

type BlockMetaData interface {
}

type BlockRewards interface {
}

type BlockGas interface {
}

type Block interface {
	RootHash() string
	Rewards() BlockRewards
	MetaData() BlockMetaData
	Number() uint64
	Hash() string
	ParentBlockHash() string
	Timestamp() uint64
	Gas() BlockGas
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	Copy() Block
}
