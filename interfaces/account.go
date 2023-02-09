package interfaces

type Account interface {
	Address() Address
	Balance() BaseToken
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	Copy() Account
}
