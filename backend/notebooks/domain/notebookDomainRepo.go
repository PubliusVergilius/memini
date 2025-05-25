package domain

import "notebooks/notebooks/dto"

type INotebookRepo interface {
	GetAllNotes() []dto.Note
	GetNoteById(id dto.ID) dto.Note
	CreateNote(note dto.Note)
	GetProfilesByUsername(username string) []dto.Profile
}
