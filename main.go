package main

import (
	"log"
	"net"

	"github.com/JairDavid/go-grpc-intro/database"
	studentserverconfig "github.com/JairDavid/go-grpc-intro/studentServerConfig"
	"github.com/JairDavid/go-grpc-intro/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository()

	if err != nil {
		log.Fatal(err)
	}

	server := studentserverconfig.NewStudentServer(repo)

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}

}
