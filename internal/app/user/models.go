package user

import "errors"

type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

var ErrDataBase = errors.New("Database connection error")
