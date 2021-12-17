package api

import (
	"bubble/models"
	"bubble/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	all, _ := models.GetAllTodo()
	fmt.Println(all)
	fmt.Println(len(all))
	if len(all) >= 20 {
		response.FailWithMessage("太多了，一天做不完的哦", c)
		return
	}
	if err := models.CreateATodo(&todo); err != nil {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func GetTodoList(c *gin.Context) {
	if todoList, err := models.GetAllTodo(); err != nil {
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(todoList, "查询成功", c)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		response.FailWithMessage("无效的id", c)
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		response.FailWithMessage("更新目标查找失败", c)
		return
	}
	_ = c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		response.FailWithMessage("删除目标查询失败", c)
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
