package group

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Test add group
func testAddGroupHandler(t *testing.T) {

	b, err := json.Marshal(map[string]string{"name": "super group"})

	if err != nil {
		t.Error(err)
		return
	}

	body := bytes.NewReader(b)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/group", body)

	if err != nil {
		t.Error(err)
		return
	}

	h := http.HandlerFunc(testService.AddGroupHandler)

	h.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Error status should be %d but got %d", http.StatusCreated, w.Code)
	}
}

// Test get group
func testGetGroupHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/group/1", nil)

	if err != nil {
		t.Error(err)
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/group/{id}", testService.GetGroupHandler)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Error status should be %d but got %d", http.StatusOK, w.Code)
	}
}
