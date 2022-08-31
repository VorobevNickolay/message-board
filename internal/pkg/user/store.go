package user

import "context"

type Store interface {
	CreateUser(ctx context.Context, name, password string) (User, error)
	FindUserById(ctx context.Context, id string) (User, error)
	FindUserByNameAndPassword(ctx context.Context, name, password string) (User, error)
	GetUsers(ctx context.Context) ([]*User, error)
}
