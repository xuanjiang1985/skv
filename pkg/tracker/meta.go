package tracker

import "time"

type Meta struct {
	Key         string     `json:"key"`
	StorageNode []string   `json:"storage_node"`
	ExpireAt    *time.Time `json:"expire_at"` // if this column equals nil, will be not expire forever.
	QueryCount  int        `json:"query_count"`
	DeletedTag  bool       `json:"deleted_tag"`
}
