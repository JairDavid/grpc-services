package repository

import (
	"context"

	"github.com/JairDavid/go-grpc-intro/domain"
)

var implementationStudent StudentRepository

type StudentRepository interface {
	GetStudent(ctx context.Context, id string) (*domain.Student, error)
	SetStudent(ctx context.Context, student *domain.Student) error
}

func NewStudentRepository(repo StudentRepository) {
	implementationStudent = repo
}

func GetStudent(ctx context.Context, id string) (*domain.Student, error) {
	return implementationStudent.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *domain.Student) error {
	return implementationStudent.SetStudent(ctx, student)
}
