package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var _ Store = (*postgresStore)(nil)

type postgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(pool *pgxpool.Pool) Store {
	return &postgresStore{pool}
}

var selectUsers = "SELECT id, name,password FROM users "

func scanUser(row pgx.Row) (User, error) {
	var u User
	err := row.Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
func scanUsers(rows pgx.Rows) ([]*User, error) {
	var users []*User
	var u User
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Username, &u.Password)
		users = append(users, createPointer(u))
		if err != nil {
			return []*User{}, fmt.Errorf("failed to select users from db %w", err)
		}
	}
	return users, nil
}

func (s *postgresStore) CreateUser(ctx context.Context, name, password string) (User, error) {
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
func (s *postgresStore) FindUserByID(ctx context.Context, id string) (User, error) {
	sql := selectUsers + "WHERE id = $1"
	row := s.pool.QueryRow(ctx, sql, id)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("failed to select user from db %w", err)
	}
	return user, nil
}

func (s *postgresStore) FindUserByName(ctx context.Context, name string) (User, error) {
	sql := selectUsers + "WHERE name = $1"
	row := s.pool.QueryRow(ctx, sql, name)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("failed to select users from db %w", err)
	}
	return user, nil
}

func (s *postgresStore) GetUsers(ctx context.Context) ([]*User, error) {
	sql := selectUsers
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return []*User{}, err
	}
	users, err := scanUsers(rows)
	if err != nil {
		return []*User{}, fmt.Errorf("failed to select users from db %w", err)
	}

	return users, nil
}
