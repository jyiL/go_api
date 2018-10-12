package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    jwtgo "github.com/dgrijalva/jwt-go"
    "time"
    "strings"
    myjwt "go_api/service"
)

func LoginApi(c *gin.Context) {
    name := c.Request.FormValue("name")
    password := c.Request.FormValue("password")

    if strings.ToLower(name) != "admin" {
        c.JSON(http.StatusPreconditionFailed, gin.H{
            "message":    "Error logging in, name error",
        })
        return
    }

    if password != "admin" {
        c.JSON(http.StatusPreconditionFailed, gin.H{
            "message":    "Error logging in, password error",
        })
        return
    }

    generateToken(c)
    return
}

// 生成token
func generateToken(c *gin.Context) {
    j := &myjwt.JWT{
        []byte("newtrekWang"),
    }
    claims := myjwt.CustomClaims{
        jwtgo.StandardClaims{
            NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
            ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
            Issuer:    "newtrekWang",                   // 签名的发行者
        },
    }

    token, err := j.CreateToken(claims)

    if err != nil {
        c.JSON(http.StatusExpectationFailed, gin.H{
            "message":    err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":    "登录成功！",
        "token":   token,
    })
    return
}