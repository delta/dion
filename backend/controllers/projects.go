package controllers

import (
	"delta.nitt.edu/dion/repository"
)

func GetAllProjects(id uint) ([]string, error) {
	return repository.GetAllProjects(id)
}

func UpdateProject(old_name string, new_name string, id uint) error {
	return repository.UpdateProject(old_name, new_name, id)
}

//	func GetProject(c *gin.Context, db *gorm.DB, logger *zap.SugaredLogger, id string) (models.Project, error) {
//		projectId, err := strconv.ParseUint(id, 10, 64)
//		if err != nil {
//			logger.Errorf("Couldn't parse configuration id: %s\n", id)
//			return models.Project{}, err
//		} else {
//			return repository.GetProject(db, logger, projectId)
//		}
//	}

func AddProject(name string, id uint) error {
	return repository.AddProject(name, id)
}

func DeleteProject(name string, user_id uint) error {
	return repository.DeleteProject(name, user_id)
}
