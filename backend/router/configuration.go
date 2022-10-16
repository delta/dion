package router

import (
	"delta.nitt.edu/dion/controllers"
	"delta.nitt.edu/dion/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func addConfigurationRoutes(configurations *gin.RouterGroup, db *gorm.DB, logger *zap.SugaredLogger) {
	configurations.GET("/", func (c *gin.Context) {
		result, err := controllers.GetAllConfigurations(c, db, logger)
		statusCode, unwrappedResult := util.UnwrapResult(&result, err)
		c.JSON(statusCode, *unwrappedResult)
	})

	configurations.GET("/:id", func (c *gin.Context) {
		result, err := controllers.GetConfiguration(c, db, logger, c.Param("id"))
		statusCode, unwrappedResult := util.UnwrapResult(&result, err)
		c.JSON(statusCode, *unwrappedResult)
	})
}