package store

type Store interface {
	UserStore() UserStore
	RoleStore() RoleStore
	UserRoleStore() UserRoleStore
	UserLoginStore() UserLoginStore
	Close()
}

type UserStore interface {
	GetUserByUsername(username string)
}

type RoleStore interface {
}

type UserRoleStore interface {
}

type UserLoginStore interface {
}
