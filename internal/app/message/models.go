package message

import "errors"

type MessageModel struct {
	username string `json:"username"`
	text     string `json:"text"`
}

var ErrDataBase = errors.New("database connection error")
