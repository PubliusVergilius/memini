package main

import (
	"fmt"
	"log"
	"net/http"
	"notebooks/notebooks"
	"strings"
)

func main() {
	port := ":5000"
	fmt.Printf("Listen on port %s", strings.TrimPrefix(port, ":"))
	server := &notebooks.NotebookServer{Store: notebooks.NewInMemoryNotebookStore()}
	log.Fatal(http.ListenAndServe(port, server))
}
