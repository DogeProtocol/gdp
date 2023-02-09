package interfaces

type TransactionType uint16

const (
	TxTypeSendBaseToken  TransactionType = 0
	TxTypeCreateContract                 = 1
	TxTypeCallContract                   = 2
)

type TransactionMetaData interface {
}

type Transaction interface {
	MetaData() TransactionMetaData
	Type() TransactionType
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	Sender() Address
	SenderNonce() uint64
	Signature() []byte
	To() Address
	Copy() Transaction
}
