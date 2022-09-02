package message

import "errors"

type Message struct {
	ID     string `json:"id"`
	UserId string `json:"userId"`
	Text   string `json:"text"`
}

var ErrMessageNotFound = errors.New("message was not found")
var ErrEmptyMessage = errors.New("empty message text")
