package storage

type Storager interface {
	Query(key []byte) ([]byte, error)
	Store(key []byte, value []byte) error
	Delete(key []byte) error
}

type StorageInstance struct {
	Node []string // eg: []string{"127.0.0.1:6211"}
}

func NewStorager(db *Storager) *Storager {
	return db
}
