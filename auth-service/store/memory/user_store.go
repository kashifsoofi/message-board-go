package memory

import "github.com/kashifsoofi/messageboard/auth/store"

type MemoryUserStore struct {
	*MemoryStore
}

func NewMemoryUserStore(memoryStore *MemoryStore) store.UserStore {
	return &MemoryUserStore{memoryStore}
}

func (s MemoryUserStore) GetUserByUsername(username string) {

}
