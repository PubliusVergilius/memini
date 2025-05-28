package database

import (
	"notebooks/notebooks/domain"
)

// Where all repos are assigned
type NotebookStore struct {
	NotebookService domain.INotebookRepo
}

func NewNotebookStore() *NotebookStore {
	// database connection argument
	return &NotebookStore{
		NotebookService: domain.NewNotebookRepoMongo(),
	}
}
