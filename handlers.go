package group

import (
	"context"
	"encoding/json"
	"net/http"
)

// AddGroupHandler add a new group
func (s *Service) AddGroupHandler(w http.ResponseWriter, r *http.Request) {

	form := GForm{}
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	id, err := s.addGroup(ctx, form)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")

	out := map[string]int64{"id": id}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// GetGroupHandler get a group
func (s *Service) GetGroupHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getURLID(r)

	if err != nil {
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	group, err := s.getGroup(ctx, id)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(group); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ListGroupsHandler list groups
func (s *Service) ListGroupsHandler(w http.ResponseWriter, r *http.Request) {

	limit, offset := getParamURLLimitOffset(r)

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	groups, err := s.listGroups(ctx, limit, offset)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(groups); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// UpdateGroupHandler update a group
func (s *Service) UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getURLID(r)

	if err != nil {
		return
	}

	form := GForm{}
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	err = s.updateGroup(ctx, id, form)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-Type", "application/json")
}

// DeleteGroupHandler delete a group
func (s *Service) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {

}

// JoinGroupHandler join a group
func (s *Service) JoinGroupHandler(w http.ResponseWriter, r *http.Request) {

}

// LeaveGroupHandler leave a group
func (s *Service) LeaveGroupHandler(w http.ResponseWriter, r *http.Request) {

}
