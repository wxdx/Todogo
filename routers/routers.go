package routers

import (
	"Todogo/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {

	//gin.SetMode(gin.ReleaseMode)
	app := fiber.New()
	app.Static("/", "template/index.html")
	app.Static("/static", "static")

	v1 := app.Group("/v1")

	v1.Post("/todo", controller.Create)
	v1.Get("/todo", controller.FindList)
	v1.Put("/todo/:id", controller.Update)
	v1.Delete("/todo/:id", controller.Delete)

	return app

}
