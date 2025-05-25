package notebooks

import (
	"io"
	"log"
	"net/http"
	"notebooks/notebooks/domain"
	"notebooks/notebooks/handlers"
	"notebooks/notebooks/service"
)

/**************** Server ******************/
// NoteStore stores notes

type NotebookServer struct {
	Store domain.INotebookRepo
	http.Handler
}

func NewNotebookServer(store domain.INotebookRepo) *NotebookServer {

	notebookServer := new(NotebookServer)
	notebookServer.Store = store

	/**************** Routing ******************/
	router := http.NewServeMux()

	notebookHandler := handlers.NewNotebookHandler(service.NewNotebookService(store))

	router.Handle("/notes", http.HandlerFunc(notebookHandler.NotebookHandler))
	router.Handle("/notes/", http.HandlerFunc(notebookHandler.NotebookHandler))
	router.Handle("/profile", http.HandlerFunc(notebookHandler.NotebookHandler))
	router.Handle("/profile/", http.HandlerFunc(notebookHandler.NotebookHandler))

	notebookServer.Handler = router

	return notebookServer
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
