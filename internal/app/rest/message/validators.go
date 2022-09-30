package message

import (
	"errors"
	"message-board/internal/app"
)

var (
	ErrTextEmpty   = errors.New("empty text")
	ErrIDEmpty     = errors.New("empty id")
	ErrUserIDEmpty = errors.New("empty userID")
)

func (r PostRequest) Validate() error {
	ve := app.NewValidationErrors()
	if len(r.Text) == 0 {
		ve.Errors["text"] = ErrTextEmpty.Error()
	}
	if len(r.UserID) == 0 {
		ve.Errors["userid"] = ErrUserIDEmpty.Error()
	}
	if len(ve.Errors) == 0 {
		return nil
	}
	return ve
}

func (r UpdateRequest) Validate() error {
	ve := app.NewValidationErrors()
	if len(r.Text) == 0 {
		ve.Errors["text"] = ErrTextEmpty.Error()
	}
	if len(r.UserID) == 0 {
		ve.Errors["userid"] = ErrUserIDEmpty.Error()
	}
	if len(r.ID) == 0 {
		ve.Errors["id"] = ErrIDEmpty.Error()
	}
	if len(ve.Errors) == 0 {
		return nil
	}
	return ve
}
