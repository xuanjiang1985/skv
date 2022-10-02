package tracker

type Tracker interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte, expireSeconds int) error
	Delete(key []byte) error
}

type TrackerClient struct {
}

func NewTrackerClient() *TrackerClient {
	return &TrackerClient{}
}

func (tc *TrackerClient) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (tc *TrackerClient) Put(key []byte, value []byte, expireSeconds int) error {
	return nil
}

func (tc *TrackerClient) Delete(key []byte) error {
	return nil
}
