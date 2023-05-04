package main

import (
	todoitem "todo-api/app/controller"
	"todo-api/config"
	"todo-api/infra/database"

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

	database.ConfigurePostgresDatabase()

	app := fiber.New()
	app.Use(
		logger.New(),
	)

	router := app.Group("/todo")
	router.Get("/swagger/*", swagger.HandlerDefault)
	todoitem.NewTodoItemController(router.Group("/item"))

	err := app.Listen(":" + config.Env.ApiPort)
	if err != nil {
		panic("could not initialize API")
	}
}
