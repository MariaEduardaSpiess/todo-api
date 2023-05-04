package service

import (
	"todo-api/app/dto"
	"todo-api/app/repository"
)

type TodoItemService struct {
	repo *repository.TodoItemRepository
}

func NewTodoItemService() *TodoItemService {
	return &TodoItemService{
		repo: repository.NewTodoItemRepository(),
	}
}

func (service *TodoItemService) CreateTodoItem(todoItem *dto.CreateTodoItemRequest) (*dto.TodoItemResponse, error) {
	created, err := service.repo.CreateTodoItem(todoItem.ToEntity())
	if err != nil {
		return nil, err
	}

	var response dto.TodoItemResponse
	return response.FromEntity(created), nil
}

func (service *TodoItemService) GetTodoItems() (*dto.GetTodoItemResponse, error) {
	items, err := service.repo.GetTodoItem()
	if err != nil {
		return nil, err
	}
	var response dto.GetTodoItemResponse
	return response.FromEntity(items), nil
}
