package interfaces

type KeyValue interface {
	Key() []byte
	Value() []byte
}

type Iterator interface {
	Prefix() []byte
	Limit() []byte
	Get() (KeyValue, error)
	Close() error
	Next() error
}

type WriteOperation interface {
}

type PutOperation struct {
	WriteOperation
	KeyValue
}

type DeleteOperation interface {
	WriteOperation
	[]byte
}

type Database interface {
	Create(fileName string) error
	Open(fileName string)
	Close() error
	Get(key []byte) (value []byte, err error)
	Put(kv KeyValue) (err error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	BatchWrite(operations []WriteOperation)
	OpenIterator() Iterator
}
