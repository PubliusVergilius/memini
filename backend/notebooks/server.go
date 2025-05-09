package notebooks

import (
	"fmt"
	"net/http"
)

type StubNotebookStore struct {
	notes map[string]string
}

func NotebooksServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá")
	
}
