package message

import "errors"

type MessageResponse struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Text   string `json:"text"`
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
