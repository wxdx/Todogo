package controller

import (
	"Todogo/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//
//func IndexHandler(c *gin.Context) {
//	c.HTML(http.StatusOK, "index.html", nil)
//}

func Create(c *fiber.Ctx) error {
	var t models.Todo
	err := c.BodyParser(&t)

	if err != nil {
		return c.JSON(err.Error())
	} else {
		err := models.Create(&t)

		if err != nil {
			fmt.Printf("exec failed, err:%v\n", err)
			return c.JSON(err.Error())
		} else {
			return c.JSON(fiber.Map{"todo": t})
		}

	}

}

func FindList(c *fiber.Ctx) error {

	todoList, err := models.FindList()

	if err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return c.JSON(fiber.Map{"error": err.Error()})
	} else {
		return c.JSON(todoList)
	}

}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")

	err := models.Update(id)

	if err != nil {
		fmt.Printf("get update id failed, err:%v\n", err)
		return c.JSON(fiber.Map{"error": err.Error()})
	} else {
		return c.JSON(id)
	}

}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	err := models.Delete(id)

	if err != nil {
		fmt.Printf("get delete id failed, err:%v\n", err)
		return c.JSON(fiber.Map{"error": err.Error()})
	} else {
		fmt.Printf("delete data success")
		return c.JSON(id)
	}

}
