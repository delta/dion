package controllers

import (
	"delta.nitt.edu/dion/models"
	"delta.nitt.edu/dion/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

func GetAllConfigurations(c *gin.Context, db *gorm.DB, logger *zap.SugaredLogger) ([]models.Configuration, error) {
	return repository.GettAllConfigurations(db, logger)
}

func GetConfiguration(c *gin.Context, db *gorm.DB, logger *zap.SugaredLogger, id string) (models.Configuration, error) {
	configurationId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Couldn't parse configuration id: %s\n", id)
		return models.Configuration{}, err
	} else {
		return repository.GetConfiguration(db, logger, configurationId)
	}
}