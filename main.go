package main

import (
	"bubble/api"
	"bubble/dao"
	"bubble/models"
	"bubble/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate go build

// set GOOS=linux
//set GOARCH=amd64
//set CGO_ENABLED=0
//go:generate set GOARCH=amd64
//go:generate set GOOS=linux
//go:generate go build main.go
func main() {
	// 创建一个默认的路由引擎
	gin.ForceConsoleColor()
	u := models.Todo{}
	fmt.Println(u)
	r := router.SetupRouter()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", api.IndexHandler)
	r.GET("/todo_list", api.IndexHandler)
	r.GET("/v1/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	dao.InitMySQL()
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务

	r.Run(":9090")
}
