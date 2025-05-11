package notebooks

import (
	"log"
	"strconv"
	"sync"
)

// NewInMemoryNotebookStore initializes an empty notebook store.
func NewInMemoryNotebookStore() *InMemoryNotebookStore {
	return &InMemoryNotebookStore{
		map[ID]Note{},
		sync.RWMutex{},
	}
}

// InMemory PlayerStore collects data about notebooks in memory.
type InMemoryNotebookStore struct {
	Notes map[ID]Note
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

// SaveNote will record a new notebook
func (i *InMemoryNotebookStore) SaveNote (note Note) {
	i.lock.Lock()
	defer i.lock.Unlock()
	// BUG: delete operation will cause an ID conflict
	id := strconv.Itoa(len(i.Notes)+1)
	log.Printf("items saved: %d", len(i.Notes))
	i.Notes[ID(id)] = note 
}

func (i *InMemoryNotebookStore) GetNoteById(id ID) Note {
	i.lock.RLock()	
	defer i.lock.RUnlock()
	return i.Notes[id]
}
