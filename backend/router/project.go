package router

import (
	"delta.nitt.edu/dion/controllers"
	"delta.nitt.edu/dion/models"
	"delta.nitt.edu/dion/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func addProjectRoutes(projects *gin.RouterGroup, db *gorm.DB, logger *zap.SugaredLogger) {
	projects.GET("/", func (c *gin.Context) {
		result, err := controllers.GetAllProjects(c, db, logger)
		statusCode, unwrappedResult := util.UnwrapResult(&result, err)
		c.JSON(statusCode, *unwrappedResult)
	})

	projects.GET("/:id", func (c *gin.Context) {
		result, err := controllers.GetProject(c, db, logger, c.Param("id"))
		statusCode, unwrappedResult := util.UnwrapResult(&result, err)
		c.JSON(statusCode, *unwrappedResult)
	})

	projects.POST("/", func (c *gin.Context) {
		result, err := controllers.AddProject(c, db, logger)
		statusCode, unwrappedResult := util.UnwrapResult(&result, err)
		c.JSON(statusCode, *unwrappedResult)
	})

	projects.DELETE("/:id", func (c *gin.Context) {
		err := controllers.DeleteProject(c, db, logger, c.Param("id"))
		result := models.Project{}
		statusCode, unwrappedResult := util.UnwrapResult(&result, err)
		c.JSON(statusCode, *unwrappedResult)
	})
}