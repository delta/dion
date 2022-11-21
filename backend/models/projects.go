package models

import "time"

type Project struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"index:idx_project_name,unique" json:"name"`
	UserId    uint   `gorm:"index:idx_project_name,unique"`
	User      User
}
