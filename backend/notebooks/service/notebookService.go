package service

import (
	"fmt"
	"net/http"
	"notebooks/notebooks/domain"
	"notebooks/notebooks/dto"
)

// Service controle de acesso ao banco
type INoteService interface {
	GetAllNotes() []dto.Note
	GetNoteById(id dto.ID) dto.Note
	CreateNote(note dto.Note)
}

type NotebookService struct {
	repo domain.INotebookRepo
}

func (n *NotebookService) GetAllNotes() []dto.Note {

	return n.repo.GetAllNotes()
}

func (n *NotebookService) GetNoteById(noteId dto.ID) dto.Note {
	/** Find note by ID */
	note := n.repo.GetNoteById(noteId)

	return note
}

func (n *NotebookService) CreateNote(w http.ResponseWriter, note dto.Note) {
	n.repo.CreateNote(note)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Note successfully added!")
}
func NewNotebookService(repo domain.INotebookRepo) NotebookService {
	return NotebookService{
		repo: repo,
	}
}
