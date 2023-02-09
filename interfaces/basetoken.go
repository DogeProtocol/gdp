package interfaces

type BaseToken interface {
	DenominationType() DenominationType
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	ToString()
	Add(tokens BaseToken) error
	Subtract(tokens BaseToken) error
	Multiply(tokens BaseToken) error
	Divide(tokens BaseToken) error
}
