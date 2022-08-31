package main

import (
	"message-board/internal/app"
	"message-board/internal/app/message"
	"message-board/internal/app/user"
	messagepkg "message-board/internal/pkg/message"
	userpkg "message-board/internal/pkg/user"
)

func main() {
	userRouter := user.NewRouter(userpkg.NewInMemoryStore())
	messageRouter := message.NewRouter(messagepkg.NewInMemoryStore())
	router := app.NewRouter(userRouter, messageRouter)
	router.SetUpRouter()
	router.Run()
}
