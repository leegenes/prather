package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leegenes/prather/config"
	"github.com/leegenes/prather/models"
	"io/ioutil"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	var dbConfig config.DbConfig
	err := config.GetConfig(&dbConfig)
	if err != nil {
		return
	}
	database, err := models.InitDb(&dbConfig)
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db: database}

	router := mux.NewRouter()
	router.HandleFunc("/notes", env.NotesHandler).Methods("GET", "POST")
	router.HandleFunc("/notes/{id}", env.NoteHandler).Methods("GET", "PUT", "DELETE")

	http.ListenAndServe(":3000", router)
}
func (env *Env) NotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//return all
		notes, err := models.GetNotes(env.db)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		for _, note := range notes {
			fmt.Fprintf(w, "%s, %s, %s\n", note.Id, note.Title, note.Text)
		}
	}

	if r.Method == http.MethodPost {
		body, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			fmt.Println(readErr)
			http.Error(w, http.StatusText(400), 400)
		}

		note := models.Note{}
		if err := json.Unmarshal(body, &note); err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}
		// create new note
		resp, err := models.CreateNote(env.db, &note)

		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		fmt.Fprintf(w, "%d", resp.Id)
	}
}

func (env *Env) NoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("worked")
}


