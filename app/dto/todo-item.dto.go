package dto

import (
	"todo-api/app/entity"
)

type CreateTodoItemRequest struct {
	Description string
}

func (body *CreateTodoItemRequest) ToEntity() *entity.TodoItemEntity {
	return &entity.TodoItemEntity{
		Description: body.Description,
	}
}

type TodoItemResponse struct {
	ID          uint
	Description string
}

func (response *TodoItemResponse) FromEntity(entity *entity.TodoItemEntity) *TodoItemResponse {
	return &TodoItemResponse{
		ID:          entity.ID,
		Description: entity.Description,
	}
}

type GetTodoItemResponse struct {
	Items []TodoItemResponse
}

func (response *GetTodoItemResponse) FromEntity(entities *[]entity.TodoItemEntity) *GetTodoItemResponse {
	var getTodoItemResponse GetTodoItemResponse
	getTodoItemResponse.Items = make([]TodoItemResponse, len(*entities))
	for i, item := range *entities {
		var todoItemResponse TodoItemResponse
		getTodoItemResponse.Items[i] = *todoItemResponse.FromEntity(&item)
	}

	return &getTodoItemResponse
}
