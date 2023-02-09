package interfaces

type PrivateKey interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}

type PublicKey interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}

type KeyPair interface {
	PrivateKey
	PublicKey
}

type SigKeyPair interface {
	KeyPair
}

type SigAlg interface {
	Sign(message []byte) ([]byte, error) //Nonce etc. should be self-contained
	Verify(signature []byte, key PublicKey) error
	GenerateKeyPair() SigKeyPair
}

type KemKeyPair interface {
	KeyPair
}

type KemAlg interface {
	Encapsulate(key PublicKey) (cipherText []byte, sharedSecret []byte, err error)
	Decapsulate(cipherText []byte, key PrivateKey) (sharedSecret []byte, err error)
}

type HashAlg interface {
	Hash(data []byte) (digest []byte, err error)
}

type SymmetricAlg interface {
	Encrypt(plain []byte) (encryptedData []byte, err error) //Nonce, IV etc. should be self-contained
	Decrypt(key PrivateKey, encryptedData []byte) (plain []byte, err error)
}
