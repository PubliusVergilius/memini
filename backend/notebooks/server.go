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
	/** Post note */
	if(r.Method == http.MethodPost) {
		log.Println("---POST accepted")
		w.WriteHeader(http.StatusAccepted)
	}

	/** Get note */
	if(r.Method == http.MethodGet) {
		/** Find note by ID */
		noteId := strings.TrimPrefix(r.URL.Path, "/notes/")
		note := n.Store.GetNoteById(noteId)

		if note == "" {
			w.WriteHeader(http.StatusNotFound)
		}
	}

	fmt.Fprint(w, "error test")

}
