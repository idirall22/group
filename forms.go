package group

// GForm structure
type GForm struct {
	Name string `json:"name"`
}

// ValidateForm check if form is valid
func (f *GForm) ValidateForm() error {

	if len(f.Name) <= MinGrpNameLen {
		return ErrorGrName
	}
	return nil
}
