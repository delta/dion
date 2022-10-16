package controllers

import (
	"delta.nitt.edu/dion/models"
	"delta.nitt.edu/dion/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

func GetAllProjects(c *gin.Context, db *gorm.DB, logger *zap.SugaredLogger) ([]models.Project, error) {
	return repository.GetAllProjects(db, logger)
}

func GetProject(c *gin.Context, db *gorm.DB, logger *zap.SugaredLogger, id string) (models.Project, error) {
	projectId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Couldn't parse configuration id: %s\n", id)
		return models.Project{}, err
	} else {
		return repository.GetProject(db, logger, projectId)
	}
}

func AddProject(c *gin.Context, db *gorm.DB, logger *zap.SugaredLogger) (uint64, error) {
	var newProject models.Project
	err := c.BindJSON(&newProject)
	if err != nil {
		logger.Errorf("Failed to bind project details\n")
		return 0, err
	} else {
		return repository.AddProject(db, logger, newProject.Name, 1)
	}
}

func DeleteProject(c *gin.Context, db *gorm.DB, logger *zap.SugaredLogger, id string) (error) {
	projectId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Couldn't parse configuration id: %s\n", id)
		return err
	} else {
		repository.DeleteProject(db, logger, projectId)
		return nil
	}
}