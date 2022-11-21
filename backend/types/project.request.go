package types

type UpdateProject struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}
