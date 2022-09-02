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

var selectUsers = "SELECT id, name,password FROM users "

func userToPointerArray(user *User) []interface{} {
	return []interface{}{&user.ID, &user.Username, &user.Password}
}

// todo: move hashing logic in service

func (s *postgresStore) CreateUser(ctx context.Context, name, password string) (User, error) {
	if name == "" || password == "" {
		return User{}, ErrEmptyPassword
	}
	password = createHash(password)
	_, err := s.findUserByName(ctx, name)
	if err == nil {
		return User{}, ErrUsedUsername
	}
	sql := "INSERT INTO users (name,password,created_at) VALUES ($1,$2,$3) RETURNING id"
	params := []interface{}{
		name,             // 1
		password,         // 2
		time.Now().UTC(), // 3
	}
	row := s.pool.QueryRow(ctx, sql, params...)
	var id string
	err = row.Scan(&id)
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
	sql := selectUsers + "WHERE id = $1"
	row := s.pool.QueryRow(ctx, sql, id)
	var user User
	err := row.Scan(userToPointerArray(&user)...)
	if err != nil {
		if err.Error() == ErrNoRows.Error() {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("failed to select user from db %w", err)
	}
	return user, nil
}

func (s *postgresStore) findUserByName(ctx context.Context, name string) (User, error) {
	sql := selectUsers + "WHERE name = $1"
	row := s.pool.QueryRow(ctx, sql, name)
	var user User
	err := row.Scan(userToPointerArray(&user)...)
	if err != nil {
		if err.Error() == ErrNoRows.Error() {
			return User{}, ErrUserNotFound
		}
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
	sql := selectUsers
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return []*User{}, err
	}
	var users []*User
	var user User
	for rows.Next() {
		err := rows.Scan(userToPointerArray(&user)...)
		users = append(users, createPointer(user))
		if err != nil {
			return []*User{}, fmt.Errorf("failed to select users from db %w", err)
		}
	}

	return users, nil
}
