package models

type Project struct {
	ID		uint64	`json:"id"`
	Name	string	`json:"name"`
	UserId	uint64	`json:"user_id"`
	User	User	`json:"user"`
}