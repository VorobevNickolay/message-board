package user

import "context"

type Store interface {
	CreateUser(ctx context.Context, name, password string) (User, error)
	FindUserByID(ctx context.Context, id string) (User, error)
	FindUserByName(ctx context.Context, name string) (User, error)
	GetUsers(ctx context.Context) ([]*User, error)
}
