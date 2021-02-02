package controller

import (
	"Todogo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Create(c *gin.Context) {
	var t models.Todo
	err := c.ShouldBind(&t)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		err := models.Create(&t)

		if err != nil {
			fmt.Printf("exec failed, err:%v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"todo": t,
			})
		}

	}

}

func FindList(c *gin.Context) {

	todoList, err := models.FindList()

	if err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}

}

func Update(c *gin.Context) {

	id, _ := c.Params.Get("id")

	err := models.Update(id)

	if err != nil {
		fmt.Printf("get update id failed, err:%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, id)
	}

}

func Delete(c *gin.Context) {

	id, _ := c.Params.Get("id")

	err := models.Delete(id)

	if err != nil {
		fmt.Printf("get delete id failed, err:%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("delete data success")
		c.JSON(http.StatusOK, id)
	}

}
