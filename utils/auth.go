package utils

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) (int, error) {
	session := sessions.Default(c)
	idInterface := session.Get("id")
	var id int
	if idInterface == nil {
		return -1, errors.New("Not logged in")
	} else {
		id = idInterface.(int)
	}
	return id, nil
}
