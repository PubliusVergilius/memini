package notebooks

import (
	"notebooks/notebooks/dto"
	"strconv"
	"sync"
)

// NewInMemoryNotebookStore initializes an empty notebook store.
func NewInMemoryNotebookStore() *InMemoryNotebookStore {
	return &InMemoryNotebookStore{
		map[dto.ID]dto.Note{},
		map[dto.ID]dto.Profile{},
		sync.RWMutex{},
	}
}

// InMemory PlayerStore collects data about notebooks in memory.
type InMemoryNotebookStore struct {
	Notes   map[dto.ID]dto.Note
	Profile map[dto.ID]dto.Profile
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

/************** Notes ****************/
func (i *InMemoryNotebookStore) GetAllNotes() []dto.Note {
	i.lock.RLock()
	defer i.lock.RUnlock()
	notes := make([]dto.Note, 0)

	for id, note := range i.Notes {
		if id != "" || note.ID != "" {
			notes = append(notes, note)
		}
	}
	return notes
}

// SaveNote will record a new notebook
func (i *InMemoryNotebookStore) CreateNote(note dto.Note) {
	i.lock.Lock()
	defer i.lock.Unlock()
	// BUG: delete operation will cause an ID conflict
	id := strconv.Itoa(len(i.Notes) + 1)
	i.Notes[dto.ID(id)] = note
}

func (i *InMemoryNotebookStore) GetNoteById(id dto.ID) dto.Note {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.Notes[id]
}

/************** Profile ****************/
func (i *InMemoryNotebookStore) GetProfilesByUsername(username string) []dto.Profile {
	i.lock.RLock()
	defer i.lock.RUnlock()

	profileTable := make([]dto.Profile, 0)
	for _, profile := range i.Profile {
		if profile.Username == username {
			profileTable = append(profileTable, profile)
		}
	}

	return profileTable
}
