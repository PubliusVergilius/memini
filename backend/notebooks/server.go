package notebooks

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// NoteStore stores notes

type NoteStore interface {
	GetNoteById(id string) string
	SaveNote(note string)
}

type NotebookServer struct {
	Store NoteStore
}

func setStatusCode (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
		case http.MethodPost:
			w.WriteHeader(http.StatusAccepted)
	}

	
}

func (n *NotebookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		n.showNote(w, r)
	case http.MethodPost:
		n.addNote(w)
	}

}

func (n *NotebookServer) showNote (w http.ResponseWriter, r *http.Request) {
	/** Find note by ID */
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")
	note := n.Store.GetNoteById(noteId)

	if note == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, "error test")

}

func (n *NotebookServer) addNote (w http.ResponseWriter) {
	log.Println("---POST accepted")
	n.Store.SaveNote("teste 1")
	w.WriteHeader(http.StatusAccepted)
}

