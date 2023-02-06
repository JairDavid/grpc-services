package repository

import (
	"context"
	"github.com/JairDavid/go-grpc-intro/pkg/domain"
)

var implementation Repository

type Repository interface {
	GetStudent(ctx context.Context, id string) (*domain.Student, error)
	SetStudent(ctx context.Context, student *domain.Student) error
	GetExam(ctx context.Context, id string) (*domain.Exam, error)
	SetExam(ctx context.Context, exam *domain.Exam) error
	SetQuestion(ctx context.Context, question *domain.Question) error
}

func NewStudentRepository(repo Repository) {
	implementation = repo
}

func GetStudent(ctx context.Context, id string) (*domain.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *domain.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetExam(ctx context.Context, id string) (*domain.Exam, error) {
	return implementation.GetExam(ctx, id)
}

func SetExam(ctx context.Context, exam *domain.Exam) error {
	return implementation.SetExam(ctx, exam)
}

func SetQuestion(ctx context.Context, question *domain.Question) error {
	return implementation.SetQuestion(ctx, question)
}
