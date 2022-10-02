package message

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"time"
)

type Message struct {
	ID        string
	UserID    string
	Text      string
	CreatedAt time.Time
}

var ErrMessageNotFound = errors.New("message was not found")
var ErrEmptyMessage = errors.New("empty message text")
var ErrNoRows = pgx.ErrNoRows

func createPointer(message Message) *Message {
	return &message
}
