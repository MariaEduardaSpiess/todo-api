package main

import (
	"todo-api/config"
	todoitem "todo-api/core/domain/todo-item"

	_ "todo-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title       Todo API
// @version     1.0
// @description API for a Todo App
// @host        localhost:9000
// @BasePath    /todo
func main() {
	config.LoadEnv()

	app := fiber.New()
	app.Use(
		logger.New(),
	)

	router := app.Group("/todo")
	router.Get("/swagger/*", swagger.HandlerDefault)
	todoitem.NewTodoItemHandler(router.Group("/item"))

	err := app.Listen(":" + config.Env.ApiPort)
	if err != nil {
		panic("could not initialize API")
	}
}
