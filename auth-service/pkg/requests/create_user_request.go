package requests

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func NewCreateUserRequest(username string, password string) CreateUserRequest {
	return CreateUserRequest{
		Id:       uuid.New(),
		Username: username,
		Password: password,
	}
}
