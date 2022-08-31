package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var _ Store = (*postgresStore)(nil)

type postgresStore struct {
	pool *pgxpool.Pool
}

func newPostgresStore(pool *pgxpool.Pool) Store {
	return &postgresStore{pool}
}

// todo: move hashing logic in service
func (s *postgresStore) CreateUser(ctx context.Context, name, password string) (User, error) {
	//s.pool.
	password = createHash(password)
	sql := "INSERT INTO users (name,password,created_at) VALUES ($1,$2,$3) RETURNING id"
	params := []interface{}{
		name,             // 1
		password,         // 2
		time.Now().UTC(), // 3
	}
	row := s.pool.QueryRow(ctx, sql, params...)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return User{}, fmt.Errorf("failed to insert user into db %w", err)
	}
	return User{
		ID:       id,
		Username: name,
		Password: createHash(password),
	}, nil
}
func (s *postgresStore) FindUserById(ctx context.Context, id string) (User, error) {
	panic("implement me")
}
func (s *postgresStore) FindUserByNameAndPassword(ctx context.Context, name, password string) (User, error) {
	panic("implement me")
}
func (s *postgresStore) GetUsers(ctx context.Context) ([]*User, error) {
	panic("implement me")
}
