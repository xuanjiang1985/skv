package storage

type BoltDBInstance struct {
	Node []string // eg: []string{"127.0.0.1:6211"}
}

func (bi *BoltDBInstance) Query(key []byte) ([]byte, error) {
	return nil, nil
}

func (bi *BoltDBInstance) Store(key []byte, value []byte) error {
	return nil
}

func (bi *BoltDBInstance) Delete(key []byte) error {
	return nil
}
