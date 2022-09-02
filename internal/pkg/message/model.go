package message

import (
	"errors"
	"github.com/jackc/pgx/v4"
)

type Message struct {
	ID     string `json:"id"`
	UserId string `json:"userId"`
	Text   string `json:"text"`
}

var ErrMessageNotFound = errors.New("message was not found")
var ErrEmptyMessage = errors.New("empty message text")
var ErrNoRows = pgx.ErrNoRows

func createPointer(message Message) *Message {
	return &message
}
