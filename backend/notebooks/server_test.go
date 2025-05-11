package notebooks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StubNotebookStore struct {
	notes map[ID]Note
}

func (s *StubNotebookStore) GetNoteById(id ID) Note {
	note := s.notes[id]
	return note
}

func (s *StubNotebookStore) SaveNote(note Note) {
	s.notes["1"] = note
}

func TestGETNotebooks (t *testing.T) {
	store := StubNotebookStore{
		map[ID]Note{
			"1" : "teste 1",
			"2" : "teste 2",
		},
	}
	server := &NotebookServer{&store}

	tests := []struct {
		testName           string
		noteId             ID
		expectedHTTPStatus int
		expectedNote       Note
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
			request := newGetNotebookRequest(tt.noteId)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStatus)
		})
	}
}

func TestPOSTNotebook (t *testing.T) {
	
	store := NewInMemoryNotebookStore()
	server := &NotebookServer{store}

	t.Run("it returns accepted on POST a save one", func(t *testing.T) {
		var note Note = "teste 1"
		request, _ := http.NewRequest(http.MethodPost, "/post", strings.NewReader(string(note)))
		request.Header.Set("Content-Type", "text/plain")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.Notes) != 1 {
			t.Fatalf("got %d notes, want %d", len(store.Notes), 1)
		}

		if store.Notes["1"] != note {
			t.Errorf("did not store correnct note: got %q want %q", store.Notes["1"], note)
		}
	})

	t.Run("it returns accepted on POST a save a second", func(t *testing.T) {
		var note Note = "teste 2"
		request, _ := http.NewRequest(http.MethodPost, "/post", strings.NewReader(string(note)))
		request.Header.Set("Content-Type", "text/plain")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.Notes) != 2 {
			t.Fatalf("got %d notes, want %d", len(store.Notes), 2)
		}

		if store.Notes["2"] != note {
			t.Errorf("did not store correnct note: got %q want %q", store.Notes["2"], note)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetNotebookRequest(id ID) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/notes/%s", id), nil)
	return req
}

func newPostNotebookRequest(_ Note) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/notes", nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertNotError (t testing.TB, err error) {
	t.Helper()
	t.Fatalf("was not told to error: %s", err.Error())
}

