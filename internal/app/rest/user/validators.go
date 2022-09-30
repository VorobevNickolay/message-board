package user

import (
	"errors"
	"message-board/internal/app"
)

var (
	ErrUsernameInvalid = errors.New("invalid username")
	ErrPasswordInvalid = errors.New("invalid password")
	ErrUsernameEmpty   = errors.New("empty username")
	ErrPasswordEmpty   = errors.New("empty password")
)

func (r LoginRequest) Validate() error {
	ve := app.NewValidationErrors()
	if len(r.Username) == 0 {
		ve.Errors["username"] = ErrUsernameEmpty.Error()
	}
	if len(r.Password) == 0 {
		ve.Errors["password"] = ErrPasswordEmpty.Error()
	}
	if len(ve.Errors) == 0 {
		return nil
	}
	return ve
}

func (r SignUpRequest) Validate() error {
	ve := app.NewValidationErrors()
	if len(r.Username) < 4 {
		ve.Errors["username"] = ErrUsernameInvalid.Error()
	}
	if len(r.Password) < 10 {
		ve.Errors["password"] = ErrPasswordInvalid.Error()
	}
	if len(ve.Errors) == 0 {
		return nil
	}
	return ve
}
