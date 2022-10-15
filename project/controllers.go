package project

import (
	"delta.nitt.edu/dion/models"
	"gorm.io/gorm"
)

func addProject(db *gorm.DB, name string, id int) error {
	project := models.Project{Name: name, UserId: id}
	result := db.Create(project)
	return result.Error
}

func getAllProjects(db *gorm.DB, id int) ([]string, error) {
	var projects []string
	result := db.Model(&models.Project{}).Where("user_id = ?", id).Select("name").Find(&projects)
	return projects, result.Error
}

func deleteProject(db *gorm.DB, id int, name string) error {
	result := db.Where("user_id = ? AND name = ?", id, name).Delete(&models.Project{})
	return result.Error
}

func changeProject(db *gorm.DB, id int, currentName string, changedName string) error {
	result := db.Model(&models.Project{}).Where("user_id = ? AND name = ?", id, currentName).Update("name", changedName)
	return result.Error
}
