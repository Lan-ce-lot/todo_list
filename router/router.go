package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	var Router = gin.Default()
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	todoRouter := TodoRouter{}
	//tmp := Router.Group("")
	todoRouter.InitTodoRouter(Router.Group(""))
	return Router
}
