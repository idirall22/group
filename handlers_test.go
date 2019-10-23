package group

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test add group
func testAddGroup(t *testing.T) {

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
