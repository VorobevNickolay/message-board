package user

import (
	"errors"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var ErrUserNotFound = errors.New("user was not found")
var ErrUsedUsername = errors.New("username already in use")
var ErrEmptyPassword = errors.New("empty password or username")
var ErrNoRows = errors.New("no rows in result set")
