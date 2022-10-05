package todoitem

type TodoItemService struct {
	todoItemRepo *TodoItemRepo
}

func NewTodoItemService() *TodoItemService {
	return &TodoItemService{
		todoItemRepo: NewTodoItemRepo(),
	}
}

func (service *TodoItemService) CreateTodoItem(todoItem *TodoItem) error {
	return service.todoItemRepo.CreateTodoItem(todoItem)
}

func (service *TodoItemService) GetTodoItems() (*[]TodoItem, error) {
	return service.todoItemRepo.GetTodoItems()
}
