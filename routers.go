package group

import (
	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Router comment endpoints
func (s *Service) Router() *mux.Router {
	r := &mux.Router{}

	r.HandleFunc("/groups", u.AuthnticateUser(s.ListGroupsHandler)).Methods("GET")
	r.HandleFunc("/groups/{id}", u.AuthnticateUser(s.GetGroupHandler)).Methods("GET")
	r.HandleFunc("/groups", u.AuthnticateUser(s.AddGroupHandler)).Methods("POST")
	r.HandleFunc("/groups/{id}", u.AuthnticateUser(s.UpdateGroupHandler)).Methods("PUT")
	r.HandleFunc("/groups/{id}", u.AuthnticateUser(s.DeleteGroupHandler)).Methods("DELETE")

	return r
}
