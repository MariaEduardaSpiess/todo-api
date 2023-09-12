package repository

import (
	"gorm.io/gorm"
	"todo-api/app/entity"
	"todo-api/infra/database"
)

type TodoItemRepository interface {
	CreateTodoItem(todoItem *entity.TodoItem) (*entity.TodoItem, error)
	GetTodoItems() (*[]entity.TodoItem, error)
}

type todoItemRepository struct {
	Db *gorm.DB
}

func NewTodoItemRepository() TodoItemRepository {
	return &todoItemRepository{
		Db: database.Db,
	}
}

func (repo *todoItemRepository) CreateTodoItem(todoItem *entity.TodoItem) (*entity.TodoItem, error) {
	err := database.Db.Create(todoItem).Error
	return todoItem, err
}

func (repo *todoItemRepository) GetTodoItems() (*[]entity.TodoItem, error) {
	var items []entity.TodoItem
	err := database.Db.Find(&items).Error
	return &items, err
}
