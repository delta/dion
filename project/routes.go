package project

import (
	"fmt"
	"net/http"

	"delta.nitt.edu/dion/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AddProject(db *gorm.DB, c *gin.Context, unlogger *zap.Logger) {
	id, err := utils.CheckAuth(c)
	var projectRequest ProjectRequest
	logger := unlogger.Sugar()
	if err != nil {
		logger.Infow("Unauthorized request")
		c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		return
	}
	err = c.BindJSON(&projectRequest)
	if err != nil {
		logger.Infow(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid request"})
		return
	}
	err = addProject(db, projectRequest.Name, id)
	if err != nil {
		logger.Errorw(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Couldn't add project"})
		return
	}
	logger.Infow(fmt.Sprintf("Project: %s created for user id: %d", projectRequest.Name, id))
	c.JSON(http.StatusOK, gin.H{"status": "Added successfully"})
}

func GetAllProjects(db *gorm.DB, c *gin.Context, unlogger *zap.Logger) {
	id, err := utils.CheckAuth(c)
	logger := unlogger.Sugar()
	if err != nil {
		logger.Infow("Unauthorized request")
		c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		return
	}
	projects, newerr := getAllProjects(db, id)
	if newerr != nil {
		logger.Errorw(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Couldn't get projects"})
		return
	}
	logger.Infow(fmt.Sprintf("All projects received successfully for user id %d", id))
	c.JSON(http.StatusOK, gin.H{"status": "Ok", "projects": projects})
}

func DeleteProject(db *gorm.DB, c *gin.Context, unlogger *zap.Logger) {
	id, err := utils.CheckAuth(c)
	var projectRequest ProjectRequest
	logger := unlogger.Sugar()
	if err != nil {
		logger.Infow("Unauthorized request")
		c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		return
	}
	err = c.BindJSON(&projectRequest)
	if err != nil {
		logger.Infow(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid request"})
		return
	}
	err = deleteProject(db, id, projectRequest.Name)
	if err != nil {
		logger.Errorw(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Couldn't delete project"})
		return
	}
	logger.Infow(fmt.Sprintf("Project: %s deleted for user id: %d", projectRequest.Name, id))
	c.JSON(http.StatusOK, gin.H{"status": "Deleted successfully"})
}

func ChangeProject(db *gorm.DB, c *gin.Context, unlogger *zap.Logger) {
	logger := unlogger.Sugar()
	id, err := utils.CheckAuth(c)
	if err != nil {
		logger.Infow("Unauthorized request")
		c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		return
	}
	var projectChangeRequest ProjectChangeRequest
	err = c.BindJSON(&projectChangeRequest)
	if err != nil {
		logger.Infow(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid request"})
		return
	}
	err = changeProject(db, id, projectChangeRequest.Name, projectChangeRequest.ChangedName)
	if err != nil {
		logger.Errorw(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Couldn't add project"})
		return
	}
	logger.Infow(fmt.Sprintf("Project: %s changed to %s", projectChangeRequest.Name, projectChangeRequest.ChangedName))
	c.JSON(http.StatusOK, gin.H{"status": "Changed successfully"})
}
