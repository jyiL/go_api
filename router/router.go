package router

import (
    "github.com/gin-gonic/gin"
    . "go_api/controller"
    . "go_api/controller/users"
    . "go_api/controller/rpc"
    . "go_api/middlewares"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    // user
    user := router.Group("/user")

    user.POST("/login", LoginApi)

    user.GET("/", AuthMiddleWare(), GetUsersApi)

    router.GET("/", AuthMiddleWare(), IndexApi)

    router.POST("/rpc", RpcApi)

    return router
}