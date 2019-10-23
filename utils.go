package group

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// get id fro url
func getURLID(r *http.Request) (int64, error) {

	idStr, ok := mux.Vars(r)["id"]

	if !ok {
		return 0, ErrorID
	}

	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return 0, ErrorID
	}

	err = validateID(id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// get limit and offset from url
func getParamURLLimitOffset(r *http.Request) (int, int) {

	limitString := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitString)

	if err != nil {
		limit = 0
	}

	offsetString := r.URL.Query().Get("offset")

	offset, err := strconv.Atoi(offsetString)

	if err != nil {
		offset = DefaultGroupOffset
	}

	return offset, limit
}
