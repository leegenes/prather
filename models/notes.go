package models

import (
	"database/sql"
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

type Note struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
}

func GetNotes(db *sql.DB) ([]*Note, error) {
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := make([]*Note, 0)

	for rows.Next() {
		note := new(Note)
		err := rows.Scan(&note.Id, &note.CreatedAt, &note.UpdatedAt, &note.Title, &note.Text)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func GetNote(db *sql.DB, id uuid.UUID) (*Note, error) {
	row := db.QueryRow("SELECT * FROM notes WHERE id = $1;", id)

	note := new(Note)
	err := row.Scan(&note.Id, &note.CreatedAt, &note.UpdatedAt, &note.Title, &note.Text)
	if err != nil {
		return nil, err
	}
	
	return note, nil
}

func CreateNote(db *sql.DB, note *Note) (*Note, error) {
	row, err := db.Query(fmt.Sprintf("INSERT INTO notes (title, text) VALUES ('%s', '%s');", note.Title, note.Text))
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Scan(&note.Id, &note.CreatedAt, &note.UpdatedAt, &note.Title, &note.Text)
	if err != nil {
		return nil, err
	}

	return note, err
}


func DeleteNote(db *sql.DB, id uuid.UUID) (error) {
	_, err := db.Exec("DELETE FROM notes WHERE id = $1;", id)
	if err != nil {
		return err
	}

	return nil
}

