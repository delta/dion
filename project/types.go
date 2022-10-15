package project

type ProjectRequest struct {
	Name string
}

type ProjectChangeRequest struct {
	Name        string
	ChangedName string
}
