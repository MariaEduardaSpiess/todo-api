package service

import (
	"todo-api/app/dto"
	"todo-api/app/repository"
)

type TodoItemService interface {
	CreateTodoItem(todoItem *dto.CreateTodoItemRequest) (*dto.TodoItemResponse, error)
	GetTodoItems() (*dto.GetTodoItemResponse, error)
}

type todoItemService struct {
	repo repository.TodoItemRepository
}

func NewTodoItemService() TodoItemService {
	return &todoItemService{
		repo: repository.NewTodoItemRepository(),
	}
}

func (service *todoItemService) CreateTodoItem(todoItem *dto.CreateTodoItemRequest) (*dto.TodoItemResponse, error) {
	created, err := service.repo.CreateTodoItem(todoItem.ToEntity())
	if err != nil {
		return nil, err
	}

	var response dto.TodoItemResponse
	return response.FromEntity(created), nil
}

func (service *todoItemService) GetTodoItems() (*dto.GetTodoItemResponse, error) {
	items, err := service.repo.GetTodoItems()
	if err != nil {
		return nil, err
	}
	var response dto.GetTodoItemResponse
	return response.FromEntity(items), nil
}
