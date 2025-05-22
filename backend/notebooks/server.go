package notebooks

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

/**************** Server ******************/
// NoteStore stores notes
type ID string

type Note struct {
	ID         ID     `json:"id"`
	Body       string `json:"body"`
	UsernameID ID     `json:"username_id"`
}

type Profile struct {
	ID       ID     `json:"id"`
	Username string `json:"username"`
}

type NoteStore interface {
	GetAllNotes() []Note
	GetNoteById(id ID) Note
	SaveNote(note Note)
	GetProfilesByUsername(username string) []Profile
}

type NotebookServer struct {
	Store NoteStore
	http.Handler
}

func NewNotebookServer(store NoteStore) *NotebookServer {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			debug.PrintStack()
		}
	}()

	notebookServer := new(NotebookServer)
	notebookServer.Store = store

	/**************** Routing ******************/
	router := http.NewServeMux()

	router.Handle("/notes", http.HandlerFunc(notebookServer.notesHandler))
	router.Handle("/notes/", http.HandlerFunc(notebookServer.notesHandler))
	router.Handle("/profile", http.HandlerFunc(notebookServer.profileHandler))
	router.Handle("/profile/", http.HandlerFunc(notebookServer.profileHandler))

	notebookServer.Handler = router

	return notebookServer
}

/**************** Handlers ******************/

/******* Notes *******/
func (n *NotebookServer) notesHandler(w http.ResponseWriter, r *http.Request) {
	noteId := strings.TrimPrefix(r.URL.Path, "/notes/")

	switch r.Method {
	case http.MethodGet:
		// Get all notes
		if r.URL.Path == "/notes" {
			n.showAllNotes(w)
			return
		}

		// Get note by ID
		if strings.TrimSpace(noteId) != "" {
			n.showNoteById(w, ID(noteId))
			return
		}
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		newNote := Note{}
		decoder.Decode(&newNote)
		n.addNote(w, newNote)
	}
}

/******* Profile *******/

func (n *NotebookServer) profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	p := n.getProfileTable()
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

/**************** Controllers ******************/

/******* Profile *******/
func (n *NotebookServer) getProfileTable() []Profile {
	profileTable := n.Store.GetProfilesByUsername("Vini")

	return profileTable
}

/******* Notebook *******/
func (n *NotebookServer) showAllNotes(w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	notes := n.Store.GetAllNotes()
	json, _ := json.Marshal(notes)

	fmt.Fprint(w, string(json))
}

func (n *NotebookServer) showNoteById(w http.ResponseWriter, noteId ID) {
	/** Find note by ID */
	note := n.Store.GetNoteById(noteId)

	if note.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Note not found!")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "")
}

func (n *NotebookServer) addNote(w http.ResponseWriter, note Note) {
	n.Store.SaveNote(note)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Note successfully added!")
}

/**************** Utils ******************/

func readBody(w http.ResponseWriter, r *http.Request) []byte {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return nil
	}

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		errorMessage := "Unable to read body"
		// Client error
		http.Error(w, errorMessage, http.StatusInternalServerError)
		// Server log error
		log.Printf("%s: %s", errorMessage, err.Error())

		return nil
	}
	defer r.Body.Close()
	return bodyBytes
}
