package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	grpc2 "message-board/internal/app/grpc/user"
	"message-board/internal/app/rest"
	"message-board/internal/app/rest/message"
	"message-board/internal/app/rest/user"
	messagepkg "message-board/internal/pkg/message"
	userpkg "message-board/internal/pkg/user"
	"net"
	"os"
)

func main() {
	databaseUrl := "postgres://vorobevna:message-board@localhost:15432/postgres"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Print(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	userStore := userpkg.NewPostgresStore(dbPool)
	go initGRPC(userStore)
	userRouter := user.NewRouter(userStore)
	messageRouter := message.NewRouter(messagepkg.NewPostgresStore(dbPool))
	router := rest.NewRouter(userRouter, messageRouter)
	router.SetUpRouter()
	router.Run()
}

func initGRPC(userStore userpkg.Store) {
	s := grpc.NewServer()
	srv := grpc2.NewServer(userStore)
	grpc2.RegisterUserServiceServer(s, srv)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
