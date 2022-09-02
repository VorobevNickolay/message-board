package user

import (
	"errors"
	"github.com/jackc/pgx/v4"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var ErrUserNotFound = errors.New("user was not found")
var ErrUsedUsername = errors.New("username already in use")
var ErrEmptyPassword = errors.New("empty password or username")
var ErrNoRows = pgx.ErrNoRows

//errors.New("no rows in result set")
