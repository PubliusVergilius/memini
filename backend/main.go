package main

import (
	"fmt"
	"log"
	"net/http"
	"notebooks/notebooks"
)

func main () {
	port := ":5000"
	fmt.Printf("Listen on port %s", port)
	handler := http.HandlerFunc(notebooks.NotebooksServer)
	log.Fatal(http.ListenAndServe(port , handler))
}
