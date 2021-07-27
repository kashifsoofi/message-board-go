package memory

import "github.com/kashifsoofi/messageboard/auth/store"

type MemoryUserLoginStore struct {
	*MemoryStore
}

func NewMemoryUserLoginStore(memoryStore *MemoryStore) store.UserLoginStore {
	return &MemoryUserLoginStore{memoryStore}
}
