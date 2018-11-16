package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
	model "go_api/models"
)

func GetUsersApi(c *gin.Context) {
	var user model.Members
	result, err := user.Users()

	if err != nil {
        c.JSON(http.StatusPreconditionFailed, gin.H{
            "message": "抱歉未找到相关信息",
        })
        return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"list": result,
	})
	return
}