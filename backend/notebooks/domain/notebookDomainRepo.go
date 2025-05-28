package domain

import "notebooks/notebooks/dto"

type INotebookRepo interface {
	GetAllNotes() []dto.Note
	GetNoteById(id dto.ID) dto.Note
	CreateNote(note dto.Note)
	GetProfilesByUsername(username string) []dto.Profile
}

type NotebookRepoMongo struct{}

func (n *NotebookRepoMongo) GetAllNotes() []dto.Note {
	return []dto.Note{{}}
}

func (n *NotebookRepoMongo) GetNoteById(id dto.ID) dto.Note {
	return dto.Note{}
}

func (n *NotebookRepoMongo) CreateNote(note dto.Note) {}

func (n *NotebookRepoMongo) GetProfilesByUsername(username string) []dto.Profile {
	return []dto.Profile{{}}
}

func NewNotebookRepoMongo() *NotebookRepoMongo {
	return &NotebookRepoMongo{}
}
