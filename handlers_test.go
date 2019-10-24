package group

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
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

	r.Header.Add("Authorization", testToken)

	h := http.HandlerFunc(u.AuthnticateUser(testService.AddGroupHandler))

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

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/group/{id}", u.AuthnticateUser(testService.GetGroupHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Error status should be %d but got %d", http.StatusOK, w.Code)
	}
}

// Test list group
func testListGroupHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/group?limit=10&offset=0", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	h := http.HandlerFunc(u.AuthnticateUser(testService.ListGroupsHandler))

	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Error status should be %d but got %d", http.StatusOK, w.Code)
	}
}

// Test list group
func testUpdateGroupHandler(t *testing.T) {

	b, err := json.Marshal(map[string]string{"name": "update super group"})

	if err != nil {
		t.Error(err)
		return
	}

	body := bytes.NewReader(b)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("PUT", "/group/1", body)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/group/{id}", u.AuthnticateUser(testService.UpdateGroupHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Errorf("Error status should be %d but got %d", http.StatusNoContent, w.Code)
	}
}

// Test delete group
func testDeleteGroupHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/group/1", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/group/{id}", u.AuthnticateUser(testService.DeleteGroupHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Errorf("Error status should be %d but got %d", http.StatusNoContent, w.Code)
	}
}

// Test join group
func testJoinGroupHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/group/join/1", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/group/join/{id}", u.AuthnticateUser(testService.JoinGroupHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Errorf("Error status should be %d but got %d", http.StatusNoContent, w.Code)
	}
}

// Test leave group
func testLeaveGroupHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/group/leave/1", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/group/leave/{id}", u.AuthnticateUser(testService.LeaveGroupHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Errorf("Error status should be %d but got %d", http.StatusNoContent, w.Code)
	}
}
