package repository

import (
	"delta.nitt.edu/dion/models"
	"gorm.io/gorm/clause"
)

func GetAllProjects(id uint) ([]string, error) {
	var projects []string
	res := db.Model(&models.Project{}).Where("user_id = ?", id).Select("name").Find(&projects)
	if res.Error != nil {
		return []string{}, res.Error
	}
	return projects, nil
}

func AddProject(name string, id uint) error {
	project := models.Project{Name: name, UserId: id}
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&project)
	return res.Error
}

func DeleteProject(name string, user_id uint) error {
	res := db.Where("user_id = ?", user_id).Where("name = ?", name).Delete(&models.Project{})
	return res.Error
}

func UpdateProject(old_name string, new_name string, user_id uint) error {
	return db.Model(&models.Project{}).Where("user_id = ?", user_id).Where("name = ?", old_name).Update("name", new_name).Error
}
