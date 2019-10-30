package group

import (
	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Router comment endpoints
func (s *Service) Router(r *mux.Router) {

	sr := r.PathPrefix("/groups").Subrouter()

	sr.HandleFunc("/", u.AuthnticateUser(s.ListGroupsHandler)).Methods("GET")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.GetGroupHandler)).Methods("GET")
	sr.HandleFunc("/", u.AuthnticateUser(s.AddGroupHandler)).Methods("POST")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.UpdateGroupHandler)).Methods("PUT")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.DeleteGroupHandler)).Methods("DELETE")
}
