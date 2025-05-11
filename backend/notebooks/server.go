package notebooks

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// NoteStore stores notes
type Note string
type ID string

type NoteStore interface {
	GetNoteById(id ID) Note
	SaveNote(note Note)
}

type NotebookServer struct {
	Store NoteStore
}

func (n *NotebookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")

	switch r.Method {
	case http.MethodGet:
		n.showNote(w, ID(noteId))
	case http.MethodPost:
		bodyBytes := readBody(w, r)

		n.addNote(w, Note(bodyBytes))
	}

}

func (n *NotebookServer) showNote (w http.ResponseWriter, noteId ID) {
	/** Find note by ID */
	note := n.Store.GetNoteById(noteId)

	if note == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, "error test")

}

func (n *NotebookServer) addNote (w http.ResponseWriter, note Note) {
	log.Println("--- POST accepted ---")
	n.Store.SaveNote(note)
	w.WriteHeader(http.StatusAccepted)
}

func readBody (w http.ResponseWriter ,r *http.Request) []byte {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)	
		return nil
	}

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		errorMessage := "Unable to read body"
		// Client error
		http.Error(w, errorMessage, http.StatusInternalServerError)
		// Server log error
		log.Printf("%s: %s", errorMessage, err.Error())

		return nil 
	}
	defer r.Body.Close()
	return bodyBytes
}
