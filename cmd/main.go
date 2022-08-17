package main

import (
	"message-board/internal/app"
	"message-board/internal/pkg/message"
	"message-board/internal/pkg/user"
)

func main() {
	router := app.NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
	router.SetUpRouter()
	router.Run()
}
