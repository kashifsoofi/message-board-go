package memory

import "github.com/kashifsoofi/messageboard/auth/store"

type MemoryUserRoleStore struct {
	*MemoryStore
}

func NewMemoryUserRoleStore(memoryStore *MemoryStore) store.UserRoleStore {
	return &MemoryUserRoleStore{memoryStore}
}
