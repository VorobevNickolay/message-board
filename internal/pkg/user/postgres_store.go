package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
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
		Password: password,
	}, nil
}
func (s *postgresStore) FindUserById(ctx context.Context, id string) (User, error) {
	sql := "SELECT id, name,password FROM users WHERE id = ($1)"
	params := []interface{}{
		id,
	}
	row := s.pool.QueryRow(ctx, sql, params...)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return User{}, fmt.Errorf("failed to select users from db %w", err)
	}
	return user, nil
}

func (s *postgresStore) findUserByName(ctx context.Context, name string) (User, error) {
	sql := "SELECT id, name,password FROM users WHERE name = ($1)"
	params := []interface{}{
		name,
	}
	row := s.pool.QueryRow(ctx, sql, params...)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return User{}, fmt.Errorf("failed to select users from db %w", err)
	}
	return user, nil
}

func (s *postgresStore) FindUserByNameAndPassword(ctx context.Context, name, password string) (User, error) {
	u, err := s.findUserByName(ctx, name)
	if err != nil {
		return User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return User{}, ErrUserNotFound
	}
	return u, nil
}
func (s *postgresStore) GetUsers(ctx context.Context) ([]*User, error) {
	sql := "SELECT id, name,password FROM users"
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return []*User{}, err
	}
	var users []*User
	var user User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		users = append(users, createPointer(user))
		if err != nil {
			return []*User{}, fmt.Errorf("failed to select users from db %w", err)
		}
	}

	return users, nil
}
