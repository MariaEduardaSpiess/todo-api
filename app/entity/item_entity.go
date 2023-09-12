package entity

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Description string
}

func (TodoItem) TableName() string {
	return "todo_item"
}
