package persistence

import (
	"context"
	"database/sql"
	"github.com/JairDavid/go-grpc-intro/pkg/domain"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/repository"
)

type StudentRepositoryImp struct {
	db *sql.DB
}

func New(conn *sql.DB) repository.StudentRepository {
	return &StudentRepositoryImp{
		db: conn,
	}
}

func (repo *StudentRepositoryImp) SetStudent(ctx context.Context, student *domain.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO student (id, name, age) VALUES ($1,$2,$3)", student.Id, student.Name, student.Age)
	if err != nil {
		return err
	}
	return nil
}

func (repo *StudentRepositoryImp) GetStudent(ctx context.Context, id string) (*domain.Student, error) {
	var studentModel domain.Student
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM student WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&studentModel.Id, &studentModel.Name, &studentModel.Age)
		if err != nil {
			return nil, err
		}
		return &studentModel, nil
	}
	return &studentModel, nil
}
