package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	grpcmessage "message-board/internal/app/grpc/message"
	grpcuser "message-board/internal/app/grpc/user"
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
	messageStore := messagepkg.NewPostgresStore(dbPool)
	go initGRPC(userStore, messageStore)
	userRouter := user.NewRouter(userStore)
	messageRouter := message.NewRouter(messageStore)
	router := rest.NewRouter(userRouter, messageRouter)
	router.SetUpRouter()
	router.Run()
}

func initGRPC(userStore userpkg.Store, messageStore messagepkg.Store) {
	s := grpc.NewServer()
	userServer := grpcuser.NewServer(userStore)
	messageServer := grpcmessage.NewServer(messageStore)
	grpcuser.RegisterUserServiceServer(s, userServer)
	grpcmessage.RegisterMessageBoardServer(s, messageServer)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
