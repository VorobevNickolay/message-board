package grpc

import (
	"errors"
	"message-board/internal/app"
)

var (
	ErrTextEmpty   = errors.New("empty text")
	ErrIDEmpty     = errors.New("empty id")
	ErrUserIDEmpty = errors.New("empty userID")
)

func (r *CreateMessageRequest) Validate() error {
	ve := app.NewValidationErrors()
	if len(r.GetText()) == 0 {
		ve.Errors["text"] = ErrTextEmpty.Error()
	}
	if len(r.GetUserId()) == 0 {
		ve.Errors["userid"] = ErrUserIDEmpty.Error()
	}
	if len(ve.Errors) == 0 {
		return nil
	}
	return ve
}

func (r *UpdateMessageRequest) Validate() error {
	ve := app.NewValidationErrors()
	if len(r.GetText()) == 0 {
		ve.Errors["text"] = ErrTextEmpty.Error()
	}
	if len(r.GetUserId()) == 0 {
		ve.Errors["userid"] = ErrUserIDEmpty.Error()
	}
	if len(r.GetId()) == 0 {
		ve.Errors["id"] = ErrIDEmpty.Error()
	}
	if len(ve.Errors) == 0 {
		return nil
	}
	return ve
}
