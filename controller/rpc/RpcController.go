package controller

import (
    "net/http"
	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
	"fmt"
	"encoding/json"
	_ "go_api/config"
    config "go_api/config"
)

type Stub struct {
	Haha      func(string) (string, error)
}

type Info struct{
    Code interface{} `json:"code"`
    Message interface{} `json:"message"`
}

func RpcApi(c *gin.Context) {
	client := rpc.NewClient(config.RpcYamlConfig.Url)
	var stub *Stub
	client.UseService(&stub)

	res, _ := stub.Haha("world")

	result := Info{}

	if err := json.Unmarshal([]byte(res), &result); err != nil {
        fmt.Println(err)
        return
	}

	fmt.Println(result.Code)
	fmt.Println(result)
	fmt.Printf("%v",result) 

    c.JSON(http.StatusOK, gin.H{
		"message": "123123",
	})

    return
}
