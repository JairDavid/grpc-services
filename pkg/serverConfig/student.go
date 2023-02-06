package serverConfig

import (
	"context"
	"github.com/JairDavid/go-grpc-intro/pkg/domain"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/repository"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/studentpb"
)

type StudentServer struct {
	repository repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *StudentServer {
	return &StudentServer{repository: repo}
}

func (s *StudentServer) GetStudent(ctx context.Context, getStudentRequest *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repository.GetStudent(ctx, getStudentRequest.GetId())
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
		Id:   getStudentRequest.GetId(),
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *StudentServer) SetStudent(ctx context.Context, studentRequest *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := &domain.Student{
		Id:   studentRequest.GetId(),
		Name: studentRequest.GetName(),
		Age:  studentRequest.GetAge(),
	}

	err := s.repository.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return &studentpb.SetStudentResponse{
		Id: studentRequest.GetId(),
	}, nil
}
