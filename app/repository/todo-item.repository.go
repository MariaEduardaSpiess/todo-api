package repository

import (
	"gorm.io/gorm"
	"todo-api/app/entity"
	"todo-api/infra/database"
)

type TodoItemRepository struct {
	Db *gorm.DB
}

func NewTodoItemRepository() *TodoItemRepository {
	return &TodoItemRepository{
		Db: database.Db,
	}
}

func (repo *TodoItemRepository) CreateTodoItem(todoItem *entity.TodoItemEntity) (*entity.TodoItemEntity, error) {
	err := database.Db.Create(todoItem).Error
	return todoItem, err
}

func (repo *TodoItemRepository) GetTodoItem() (*[]entity.TodoItemEntity, error) {
	var items []entity.TodoItemEntity
	err := database.Db.Find(&items).Error
	return &items, err
}
