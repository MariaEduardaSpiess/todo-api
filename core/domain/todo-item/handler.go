package todoitem

import "github.com/gofiber/fiber/v2"

type TodoItemHandler struct {
	todoItemService *TodoItemService
}

func NewTodoItemHandler(router fiber.Router) {
	handler := TodoItemHandler{
		todoItemService: NewTodoItemService(),
	}

	router.Post("/", handler.CreateTodoItem)
	router.Get("/", handler.GetTodoItems)
}

// @Summary     Create todo item
// @Description Create todo item in the database
// @Tags        todo-item
// @Produce     json
// @Param       body    body     TodoItem true "body"
// @Success     201     {object} nil
// @Failure     400,500 {object} error
// @Router      /item [post]
func (handler TodoItemHandler) CreateTodoItem(ctx *fiber.Ctx) error {
	body := new(TodoItem)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := handler.todoItemService.CreateTodoItem(body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

// @Summary     Get todo items
// @Description Get todo items from the database
// @Tags        todo-item
// @Produce     json
// @Success     200 {object} []TodoItem
// @Failure     500 {object} error
// @Router      /item [get]
func (handler TodoItemHandler) GetTodoItems(ctx *fiber.Ctx) error {
	response, err := handler.todoItemService.GetTodoItems()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(response)
}
