package persistence

import (
	"context"
	"database/sql"
	"github.com/JairDavid/go-grpc-intro/pkg/domain"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/repository"
)

type RepositoryImp struct {
	db *sql.DB
}

func New(conn *sql.DB) repository.Repository {
	return &RepositoryImp{
		db: conn,
	}
}

func (repo *RepositoryImp) SetStudent(ctx context.Context, student *domain.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO student (id, name, age) VALUES ($1,$2,$3)", student.Id, student.Name, student.Age)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryImp) GetStudent(ctx context.Context, id string) (*domain.Student, error) {
	var studentModel domain.Student
	err := repo.db.QueryRowContext(ctx, "SELECT * FROM student WHERE id = ?", id).Scan(studentModel.Id, studentModel.Name, studentModel.Age)
	if err != nil {
		return nil, err
	}

	return &studentModel, nil

}

func (repo *RepositoryImp) GetExam(ctx context.Context, id string) (*domain.Exam, error) {
	var examModel domain.Exam
	err := repo.db.QueryRowContext(ctx, "SELECT  * FROM exam WHERE id = ?", id).Scan(examModel.Id, examModel.Name)
	if err != nil {
		return nil, err
	}

	return &examModel, nil
}

func (repo *RepositoryImp) SetExam(ctx context.Context, exam *domain.Exam) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO exam (id, name) VALUES ($1,$2)", exam.Id, exam.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryImp) SetQuestion(ctx context.Context, question *domain.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO question (id, answer,question,exam_id) VALUES ($1,$2,$3,$4)", question.Id, question.Answer, question.Question, question.ExamId)
	if err != nil {
		return err
	}
	return nil
}
