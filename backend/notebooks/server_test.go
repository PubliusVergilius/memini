package notebooks

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)



func TestGETNotebooks (t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	
	NotebooksServer(response, request)
	
	t.Run("returns notebooks", func(t *testing.T){
		got := response.Body.String()
		want := "Olá"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	
}
