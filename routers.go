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
	sr.HandleFunc("/join/{id}", u.AuthnticateUser(s.JoinGroupHandler)).Methods("POST")
	sr.HandleFunc("/leave/{id}", u.AuthnticateUser(s.LeaveGroupHandler)).Methods("POST")
}
