package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex   *sync.Mutex
	storage *map[string]Record
}

type Record struct {
	Expiration *int
	StoredAt   time.Time
	Value      any
}

func (s *Cache) Get(key string) (any, bool) {
	record, ok := (*s.storage)[key]
	if !ok {
		return nil, false
	}

	if record.Expiration != nil {
		if int(time.Now().UnixMilli()-record.StoredAt.UnixMilli()) > *record.Expiration {
			return nil, false
		}
	}

	return record.Value, ok
}

func (s *Cache) Store(key string, record Record) {
	s.mutex.Lock()

	record.StoredAt = time.Now()

	(*s.storage)[key] = record

	s.mutex.Unlock()
}

func NewCache() *Cache {
	storage := make(map[string]Record)

	return &Cache{
		mutex:   &sync.Mutex{},
		storage: &storage,
	}
}
