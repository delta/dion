package routes

import (
	"net/http"

	"delta.nitt.edu/dion/controllers"
	"delta.nitt.edu/dion/middleware"
	"delta.nitt.edu/dion/models"
	"delta.nitt.edu/dion/services/logging"
	"delta.nitt.edu/dion/types"
	"github.com/gin-gonic/gin"
)

func getUserProjects(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user := userInterface.(*models.User)
	projects, err := controllers.GetAllProjects(user.ID)
	if err != nil {
		logging.Sugared().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"projects": []models.Project{}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}

func addProject(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user := userInterface.(*models.User)
	var newProject models.Project
	err := ctx.BindJSON(&newProject)
	if err != nil {
		logging.Sugared().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Invalid input provided"})
		return
	}
	err = controllers.AddProject(newProject.Name, user.ID)
	if err != nil {
		logging.Sugared().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Couldn't add project"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Project added successfully"})
}

func deleteProject(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user := userInterface.(*models.User)
	var project models.Project
	err := ctx.BindJSON(&project)
	if err != nil {
		logging.Sugared().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Invalid input provided"})
		return
	}
	err = controllers.DeleteProject(project.Name, user.ID)
	if err != nil {
		logging.Sugared().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Couldn't delete project"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Project deleted successfully"})
}

func updateProject(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user := userInterface.(*models.User)
	var updateProject types.UpdateProject
	err := ctx.BindJSON(&updateProject)
	if err != nil {
		logging.Sugared().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Invalid input provided"})
		return
	}
	err = controllers.UpdateProject(updateProject.OldName, updateProject.NewName, user.ID)
	if err != nil {
		logging.Sugared().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Couldn't add project"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Project added successfully"})
}

func initProject() {
	group := "project"
	RouteMap[group] = RouteGroup{
		GlobalMiddleware: gin.HandlersChain{middleware.CheckAuth},
		Routes: Routes{
			{
				"GetProject",
				http.MethodGet,
				"/all",
				getUserProjects,
				nil,
			},
			{
				"AddProject",
				http.MethodPost,
				"/add",
				addProject,
				nil,
			},
			{
				"DeleteProject",
				http.MethodDelete,
				"/delete",
				deleteProject,
				nil,
			},
			{
				"UpdateProject",
				http.MethodPatch,
				"/update",
				updateProject,
				nil,
			},
		},
	}
}
