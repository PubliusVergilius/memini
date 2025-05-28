package main

import (
	"fmt"
	"log"
	"net/http"
	"notebooks/notebooks"
	"notebooks/notebooks/database"
	"strings"
)

func main() {
	port := ":5001"
	fmt.Printf("Listen on port %s", strings.TrimPrefix(port, ":"))

	fmt.Println("Connection setup is called")
	_, client, context, cancel := database.SetupMongoDB()

	defer database.CloseConnection(client, context, cancel)

	store := database.NewNotebookStore()

	server := notebooks.NewNotebookServer(store.NotebookService)
	log.Fatal(http.ListenAndServe(port, server))
}
