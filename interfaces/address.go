package interfaces

type Address interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) (Address, error)
	Copy() Address
}
