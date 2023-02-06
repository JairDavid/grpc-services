package main

import (
	"github.com/JairDavid/go-grpc-intro/database"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/exampb"
	"github.com/JairDavid/go-grpc-intro/pkg/infrastructure/persistence"
	serverconfig "github.com/JairDavid/go-grpc-intro/pkg/serverConfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatal(err)
	}

	database.Connect()

	server := serverconfig.NewExamServer(persistence.New(database.GetConnection()))

	s := grpc.NewServer()
	exampb.RegisterExamServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}

}
