package middlewares

import (
    "net/http"
    "github.com/gin-gonic/gin"
    myjwt "go_api/service"
    "strings"
)

func AuthMiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("Authorization")
        token = strings.Replace(token, "Bearer ", "", -1)
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message":    "请求未携带token，无权限访问",
            })
            c.Abort()
            return
        }

        j := myjwt.NewJWT()
        // parseToken 解析token包含的信息
        _, err := j.ParseToken(token)
        if err != nil {
            if err == myjwt.TokenExpired {
                c.JSON(http.StatusUnauthorized, gin.H{
                    "message":    "授权已过期",
                })
                c.Abort()
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{
                "message":    err.Error(),
            })
            c.Abort()
            return
        }

        c.Next()
        return
    }
}