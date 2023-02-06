package main

import (
	"github.com/JairDavid/go-grpc-intro/pkg/domain/studentpb"
	"github.com/JairDavid/go-grpc-intro/pkg/infrastructure/persistence"
	serverconfig "github.com/JairDavid/go-grpc-intro/pkg/serverConfig"
	"log"
	"net"

	"github.com/JairDavid/go-grpc-intro/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}

	database.Connect()

	server := serverconfig.NewStudentServer(persistence.New(database.GetConnection()))

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}

}
