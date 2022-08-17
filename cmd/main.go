package main

import (
	"message-board/internal/app"
	"message-board/internal/pkg/message"
)

func main() {
	router := app.NewRouter(message.NewInMemoryStore())
	router.SetUpRouter()
	router.Run()
}
