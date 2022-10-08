package tracker

type MetaData struct {
	Key         string   `json:"key"`
	StorageNode []string `json:"storage_node"`
	ExpireAt    uint64   `json:"expire_at"` // milliseconds, if this column equals 0, will be not expire forever.
}

type Entry struct {
	Key      []byte `json:"key"`
	Value    []byte `json:"value"`
	ExpireAt uint64 `json:"expire_at"`
}
