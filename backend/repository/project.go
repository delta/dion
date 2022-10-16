package repository

import (
	"delta.nitt.edu/dion/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetAllProjects(db *gorm.DB, logger *zap.SugaredLogger) ([]models.Project, error) {
	projects := []models.Project{}
	logger.Debugf("Fetching all projects")
	result := db.Find(&projects)
	if result.Error != nil {
		logger.Errorf("Couldn't fetch projects: %s\n", result.Error)
	}
	return projects, result.Error
}

func GetProject(db *gorm.DB, logger *zap.SugaredLogger, id uint64) (models.Project, error) {
	project := models.Project{}
	logger.Debugf("Fetching project %d\n", id)
	result := db.Preload("User").Find(&project, id)
	if result.Error != nil {
		logger.Errorf("Couldn't fetch project with ID %d: %s\n", id, result.Error)
	}
	return project, result.Error
}

func AddProject(db *gorm.DB, logger *zap.SugaredLogger, name string, userId uint64) (uint64, error) {
	project := models.Project{Name: name, UserId: userId}
	logger.Debugf("Creating new project %s by user %d\n", name, userId)
	result := db.Create(&project)
	if result.Error != nil {
		logger.Errorf("Couldn't create new project %s by user %d\n", name, userId)
	}
	return project.ID, result.Error
}

func DeleteProject(db *gorm.DB, logger *zap.SugaredLogger, id uint64) {
	logger.Debugf("Deleting project %d\n", id)
	db.Delete(&models.Project{}, id)
}