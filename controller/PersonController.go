package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
    c.String(http.StatusOK, "It works")
}