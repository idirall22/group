package group

import "errors"

// MinGrpNameLen the minimum length for a group name
const MinGrpNameLen = 6

// GForm structure
type GForm struct {
	Name string `json:"name"`
}

// ValidateForm check if form is valid
func (f *GForm) ValidateForm() error {

	if len(f.Name) <= MinGrpNameLen {
		return errors.New("Group name length should have a minimum of 6 chars")
	}
	return nil
}
