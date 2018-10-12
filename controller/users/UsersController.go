package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetUsersApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "123123123",
	})
	return
}