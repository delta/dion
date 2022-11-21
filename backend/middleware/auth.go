package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"delta.nitt.edu/dion/controllers"
	"github.com/gin-contrib/sessions"
)

func CheckAuth(ctx *gin.Context) {
	session := sessions.Default(ctx)
	email := session.Get("email")
	if email == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Please login"})
		ctx.Abort()
		return
	} else {
		user, err := controllers.GetUser(email.(string))
		if err != nil {
			session.Delete("email")
			session.Save()
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "User not found"})
			ctx.Abort()
			return
		}
		ctx.Set("user", &user)
		ctx.Next()
	}
}
