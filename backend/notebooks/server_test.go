package notebooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETNotebooks(t *testing.T) {
	t.Run("get note by it's id from in memory notebook store", func(t *testing.T) {
		store := NewInMemoryNotebookStore()
		server := NewNotebookServer(store)
		store.Notes["1"] = Note{ID: "1", Body: "teste 1", UsernameID: "1"}
		store.Notes["2"] = Note{ID: "2", Body: "teste 2", UsernameID: "2"}

		tests := []struct {
			testName           string
			noteId             ID
			expectedHTTPStatus int
			expectedNote       Note
		}{
			{
				testName:           "Returns first note",
				noteId:             "1",
				expectedHTTPStatus: http.StatusOK,
				expectedNote:       Note{ID: "1", Body: "teste 1", UsernameID: "1"},
			},
			{
				testName:           "Returns second note",
				noteId:             "2",
				expectedHTTPStatus: http.StatusOK,
				expectedNote:       Note{ID: "2", Body: "teste 2", UsernameID: "2"},
			},
			{
				testName:           "Returns 404 on misssing note",
				noteId:             "3",
				expectedHTTPStatus: http.StatusNotFound,
				expectedNote:       Note{},
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
	})

	t.Run("get all in JSON from in memory store", func(t *testing.T) {
		store := NewInMemoryNotebookStore()
		store.Notes["1"] = Note{ID: "1", Body: "teste 1", UsernameID: "1"}
		store.Notes["2"] = Note{ID: "2", Body: "teste 2", UsernameID: "1"}

		server := NewNotebookServer(store)

		request, _ := http.NewRequest("GET", "/notes", nil)
		response := httptest.NewRecorder()
		response.Header().Set("Content-Type", "application/json")

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		want, _ := json.Marshal([]Note{
			{
				ID:         "1",
				Body:       "teste 1",
				UsernameID: "1",
			},
			{
				ID:         "2",
				Body:       "teste 2",
				UsernameID: "1",
			},
		})

		assertResponseBody(t, response.Body.String(), string(want))
	})
}

func TestPOSTNewNote(t *testing.T) {

	store := NewInMemoryNotebookStore()
	server := NewNotebookServer(store)

	t.Run("Store a new note on /notes and return accepted status", func(t *testing.T) {
		note := Note{ID: "1", Body: "teste 1", UsernameID: "1"}

		request := storeNewNote(t, note)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.Notes) != 1 {
			t.Fatalf("got %d notes, want 1", len(store.Notes))
		}

		if !reflect.DeepEqual(store.Notes["1"], note) {
			t.Fatalf("did not store the correct note got %q, want %q", store.Notes["1"], note)
		}

	})

}

func TestGETProfile(t *testing.T) {
	store := NewInMemoryNotebookStore()
	server := NewNotebookServer(store)
	store.Profile["1"] = Profile{ID: "1", Username: "Vini"}

	t.Run("it returnts 200 on /profile/", func(t *testing.T) {
		request, _ := http.NewRequest("get", "/profile", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getProfileFromResponse(t, response.Body)
		want := []Profile{
			{
				ID:       "1",
				Username: "Vini",
			},
		}

		assertStatus(t, response.Code, http.StatusOK)
		assertProfiles(t, got, want)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func storeNewNote(t *testing.T, note Note) *http.Request {
	jsonNote, err := json.Marshal(note)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v\n", err)
	}

	reqBody := bytes.NewBuffer(jsonNote)

	request, _ := http.NewRequest(http.MethodPost, "/notes", reqBody)
	request.Header.Set("Content-Type", "application/json")

	return request
}

func getProfileFromResponse(t testing.TB, body io.Reader) (profile []Profile) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&profile)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertProfiles(t testing.TB, got, want []Profile) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
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

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("response body is wrong, got %q want %q", got, want)
	}
}

func assertNotError(t testing.TB, err error) {
	t.Helper()
	t.Fatalf("was not told to error: %s", err.Error())
}
