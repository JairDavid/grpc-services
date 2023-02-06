package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var conn *sql.DB

func Connect() {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/school?sslmode=disable")
	if err != nil {
		panic(err)
	}
	conn = db
}

func GetConnection() *sql.DB {
	return conn
}
