package router

import (
	"bubble/api"
	"github.com/gin-gonic/gin"
)

type TodoRouter struct {
}

func (s *TodoRouter) InitTodoRouter(Router *gin.RouterGroup) {
	todoRputer := Router.Group("v1/")
	{
		todoRputer.POST("todo", api.CreateTodo)
		todoRputer.DELETE("todo/:id", api.DeleteATodo)
		todoRputer.PUT("todo/:id", api.UpdateTodo)
		todoRputer.GET("todo", api.GetTodoList)
	}
}
