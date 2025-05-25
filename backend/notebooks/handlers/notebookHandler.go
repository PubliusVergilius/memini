package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notebooks/notebooks/dto"
	"notebooks/notebooks/service"
	"strings"
)

type NotebookHandler struct {
	notebookService service.NotebookService
}

func (n *NotebookHandler) NotebookHandler(w http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")

	switch r.Method {
	case http.MethodGet:
		// Get all notes
		if r.URL.Path == "/notes" {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			notes := n.notebookService.GetAllNotes()

			json, _ := json.Marshal(notes)

			fmt.Fprint(w, string(json))
			return
		}

		// Get note by ID
		if strings.TrimSpace(noteId) != "" {

			if noteId == "" {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "Note not found!")
				return
			}
			note := n.notebookService.GetNoteById(dto.ID(noteId))
			json, _ := json.Marshal(note)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, json)
			return
		}

	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		newNote := dto.Note{}
		decoder.Decode(&newNote)
		n.notebookService.CreateNote(w, newNote)
	}
}

func NewNotebookHandler(service service.NotebookService) NotebookHandler {
	return NotebookHandler{
		notebookService: service,
	}
}
