package memory

import "github.com/kashifsoofi/messageboard/auth/store"

type MemoryRoleStore struct {
	*MemoryStore
}

func NewMemoryRoleStore(memoryStore *MemoryStore) store.RoleStore {
	return &MemoryRoleStore{memoryStore}
}
