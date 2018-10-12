package main

import (
    "github.com/gin-gonic/gin"
    . "go_api/controller"
    . "go_api/controller/users"
    . "go_api/middlewares"
)

func initRouter() *gin.Engine {
    router := gin.Default()

    // user
    user := router.Group("/user")

    user.POST("/login", LoginApi)

    user.GET("/", AuthMiddleWare(), GetUsersApi)

    router.GET("/", AuthMiddleWare(), IndexApi)

    // router.GET("/bank_log", AuthMiddleWare(), GetBankLogApi)

    return router
}