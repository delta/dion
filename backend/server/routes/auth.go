package routes

import (
	"net/http"

	"delta.nitt.edu/dion/controllers"
	"delta.nitt.edu/dion/middleware"
	"delta.nitt.edu/dion/models"
	"delta.nitt.edu/dion/services/logging"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user := userInterface.(*models.User)
	ctx.JSON(http.StatusOK, gin.H{"email": user.Email})
}

func callBack(ctx *gin.Context) {
	code := ctx.Query("code")
	email, err := controllers.HandleCallBack(code)
	if err != nil {
		logging.Sugared().Error(err.Error())
	}
	session := sessions.Default(ctx)
	session.Set("email", email)
	err = session.Save()
	ctx.JSON(http.StatusOK, gin.H{"email": email})
}

func initAuth() {
	group := "auth"
	RouteMap[group] = RouteGroup{
		Routes: Routes{
			{
				"GetUser",
				http.MethodGet,
				"/user",
				getUser,
				gin.HandlersChain{middleware.CheckAuth},
			},
			{
				"CallBack",
				http.MethodGet,
				"/callback",
				callBack,
				nil,
			},
		},
	}
}
