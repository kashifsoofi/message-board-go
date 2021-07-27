package memory

import (
	"github.com/kashifsoofi/messageboard/auth/store"
)

type MemoryStore struct {
	userStore      store.UserStore
	roleStore      store.RoleStore
	userRoleStore  store.UserRoleStore
	userLoginStore store.UserLoginStore
}

func NewMemoryStore() store.Store {
	memoryStore := &MemoryStore{}

	memoryStore.userStore = NewMemoryUserStore(memoryStore)

	return memoryStore
}

func (ms *MemoryStore) UserStore() store.UserStore {
	return ms.userStore
}

func (ms *MemoryStore) RoleStore() store.RoleStore {
	return ms.roleStore
}

func (ms *MemoryStore) UserRoleStore() store.UserRoleStore {
	return ms.userRoleStore
}

func (ms *MemoryStore) UserLoginStore() store.UserLoginStore {
	return ms.userLoginStore
}

func (ms *MemoryStore) Close() {
}
