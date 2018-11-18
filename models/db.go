package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/leegenes/prather/config"
)

func InitDb(config *config.DbConfig) (*sql.DB, error) {
	
	database, err := sql.Open("postgres", "user=haugenlee dbname=prather sslmode=disable")

	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
