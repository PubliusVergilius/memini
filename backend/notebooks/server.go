package notebooks

import (
	"fmt"
	"net/http"
	"strings"
)

// NoteStore stores notes
type NoteStore interface {
	GetNoteById(id string) string
}

type NotebookServer struct {
	store NoteStore
}

func (n *NotebookServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")
	note := n.store.GetNoteById(noteId)

	if note == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, "error test")
}
