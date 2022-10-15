package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name   string
	UserId int
	User   User
}
