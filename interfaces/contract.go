package interfaces

type Contract interface {
	Address() Address
	Balance() BaseToken
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}
