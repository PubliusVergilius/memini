package notebooks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubNotebookStore struct {
	notes map[string]string
}

func (s *StubNotebookStore) GetNoteById(id string) string {
	note := s.notes[id]
	return note
}

func TestGETNotebooks (t *testing.T) {
	store := StubNotebookStore{
		map[string]string{
			"1" : "teste 1",
			"2" : "teste 2",
		},
	}
	server := &NotebookServer{&store}

	tests := []struct {
		testName           string
		noteId             string
		expectedHTTPStatus int
		expectedNote       string
	}{
		{
			testName: "Returns first note",
			noteId: "1",
			expectedHTTPStatus: http.StatusOK,
			expectedNote: "teste 1",
		},
		{
			testName: "Returns second note",
			noteId: "2",
			expectedHTTPStatus: http.StatusOK,
			expectedNote: "teste 2",
		},
		{
			testName: "Returns 404 on misssing note",
			noteId: "3",
			expectedHTTPStatus: http.StatusNotFound,
			expectedNote: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			request := newGetNoteRequest(tt.noteId)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStatus)
		})
	}
}

func TestPOSTNotebook (t *testing.T) {
	store := StubNotebookStore{
		map[string]string{},
	}
	server := &NotebookServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/post", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetNoteRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/notes/%s", id), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

