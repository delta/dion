package repository

import (
	"delta.nitt.edu/dion/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GettAllConfigurations(db *gorm.DB, logger *zap.SugaredLogger) ([]models.Configuration, error) {
	configurations := []models.Configuration{}
	logger.Debugf("Fetching all configurations")
	result := db.Find(&configurations)
	if result.Error != nil {
		logger.Errorf("Couldn't fetch configurations: %s\n", result.Error)
	}
	return configurations, result.Error
}

func GetConfiguration(db *gorm.DB, logger *zap.SugaredLogger, id uint64) (models.Configuration, error) {
	configuration := models.Configuration{}
	result := db.Find(&configuration, id)
	if result.Error != nil {
		logger.Errorf("Couldn't fetch configuration with ID %d: %s\n", id, result.Error)
	}
	return configuration, result.Error
}