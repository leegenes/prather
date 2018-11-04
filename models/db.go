package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InitDb(dbName string) (*sql.DB, error) {
	database, err := sql.Open("postgres", "user=haugenlee dbname=prather sslmode=disable")

	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
