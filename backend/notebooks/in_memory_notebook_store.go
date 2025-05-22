package notebooks

import (
	"strconv"
	"sync"
)

// NewInMemoryNotebookStore initializes an empty notebook store.
func NewInMemoryNotebookStore() *InMemoryNotebookStore {
	return &InMemoryNotebookStore{
		map[ID]Note{},
		map[ID]Profile{},
		sync.RWMutex{},
	}
}

// InMemory PlayerStore collects data about notebooks in memory.
type InMemoryNotebookStore struct {
	Notes   map[ID]Note
	Profile map[ID]Profile
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

/************** Notes ****************/
func (i *InMemoryNotebookStore) GetAllNotes() []Note {
	i.lock.RLock()
	defer i.lock.RUnlock()
	notes := make([]Note, 0)

	for id, note := range i.Notes {
		if id != "" || note.ID != "" {
			notes = append(notes, note)
		}
	}
	return notes
}

// SaveNote will record a new notebook
func (i *InMemoryNotebookStore) SaveNote(note Note) {
	i.lock.Lock()
	defer i.lock.Unlock()
	// BUG: delete operation will cause an ID conflict
	id := strconv.Itoa(len(i.Notes) + 1)
	i.Notes[ID(id)] = note
}

func (i *InMemoryNotebookStore) GetNoteById(id ID) Note {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.Notes[id]
}

/************** Profile ****************/
func (i *InMemoryNotebookStore) GetProfilesByUsername(username string) []Profile {
	i.lock.RLock()
	defer i.lock.RUnlock()

	profileTable := make([]Profile, 0)
	for _, profile := range i.Profile {
		if profile.Username == username {
			profileTable = append(profileTable, profile)
		}
	}

	return profileTable
}
