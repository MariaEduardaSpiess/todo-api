package entity

import "gorm.io/gorm"

type TodoItemEntity struct {
	gorm.Model
	Description string
}

func (TodoItemEntity) TableName() string {
	return "todo_item"
}
