package database

import (
	"context"
	"database/sql"

	"github.com/JairDavid/go-grpc-intro/domain"
	_ "github.com/lib/pq"
)

type StudentDatabaseRepository struct {
	db *sql.DB
}

func NewPostgresRepository() (*StudentDatabaseRepository, error) {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/students?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &StudentDatabaseRepository{db: db}, nil
}

func (repo *StudentDatabaseRepository) SetStudent(ctx context.Context, student *domain.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO student (id, name, age) VALUES ($1,$2,$3)", student.Id, student.Name, student.Age)
	if err != nil {
		return err
	}
	return nil
}

func (repo *StudentDatabaseRepository) GetStudent(ctx context.Context, id string) (*domain.Student, error) {
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
