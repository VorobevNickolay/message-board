package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"message-board/internal/app"
	"message-board/internal/app/message"
	"message-board/internal/app/user"
	messagepkg "message-board/internal/pkg/message"
	userpkg "message-board/internal/pkg/user"
	"os"
)

func main() {
	databaseUrl := "postgres://vorobevna:message-board@localhost:15432/postgres"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Print(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	userRouter := user.NewRouter(userpkg.NewPostgresStore(dbPool))
	messageRouter := message.NewRouter(messagepkg.NewPostgresStore(dbPool))
	router := app.NewRouter(userRouter, messageRouter)
	router.SetUpRouter()
	router.Run()
}
