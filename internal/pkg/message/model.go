package message

import (
	"errors"
	"github.com/jackc/pgx/v4"
)

type Message struct {
	ID     string
	UserID string
	Text   string
}

var ErrMessageNotFound = errors.New("message was not found")
var ErrEmptyMessage = errors.New("empty message text")
var ErrNoRows = pgx.ErrNoRows

func createPointer(message Message) *Message {
	return &message
}
