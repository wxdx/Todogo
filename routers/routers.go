package routers

import (
	"Todogo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	//告诉gin 静态文件在哪个文件夹
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	//待办事项
	//添加
	v1Group.POST("/todo", controller.Create)
	//查看所有待办事项
	v1Group.GET("/todo", controller.FindList)
	//修改某个待办事项
	v1Group.PUT("/todo/:id", controller.Update)
	//删除某个待办事项
	v1Group.DELETE("/todo/:id", controller.Delete)

	return r

}
