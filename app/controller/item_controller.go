package controller

import (
	"github.com/gofiber/fiber/v2"
	"todo-api/app/dto"
	"todo-api/app/service"
)

type TodoItemController struct {
	service service.TodoItemService
}

func NewTodoItemController(router fiber.Router) {
	handler := TodoItemController{
		service: service.NewTodoItemService(),
	}

	router.Post("/", handler.CreateTodoItem)
	router.Get("/", handler.GetTodoItems)
}

// @Summary     CreateTodoItem todo item
// @Description CreateTodoItem todo item in the database
// @Tags        todo-item
// @Produce     json
// @Param       body    body     CreateTodoItemRequest true "body"
// @Success     201     {object} nil
// @Failure     400,500 {object} error
// @Router      /item [post]
func (controller TodoItemController) CreateTodoItem(ctx *fiber.Ctx) error {
	body := new(dto.CreateTodoItemRequest)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	created, err := controller.service.CreateTodoItem(body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(created)
}

// @Summary     Get todo items
// @Description Get todo items from the database
// @Tags        todo-item
// @Produce     json
// @Success     200 {object} []CreateTodoItemRequest
// @Failure     500 {object} error
// @Router      /item [get]
func (controller TodoItemController) GetTodoItems(ctx *fiber.Ctx) error {
	response, err := controller.service.GetTodoItems()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(response)
}
