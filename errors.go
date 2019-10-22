package group

import "errors"

var (
	// ErrorID when id is not valid
	ErrorID = errors.New("Id not valid")

	// ErrorGrName when group name length is less than MinGrpNameLen
	ErrorGrName = errors.New("Group name length should have a minimum of 6 chars")
)
