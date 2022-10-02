package message

import (
	"errors"
	"time"
)

type MessageResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created-at"`
}

type PostRequest struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type UpdateRequest struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

var ErrDataBase = errors.New("database connection error")
