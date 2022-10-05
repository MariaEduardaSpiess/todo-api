package todoitem

import (
	"todo-api/config"
	"todo-api/infra/database"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type TodoItemRepo struct {
	TableName string
}

func NewTodoItemRepo() *TodoItemRepo {
	return &TodoItemRepo{
		TableName: config.Env.Database.TableName,
	}
}

func (repo *TodoItemRepo) CreateTodoItem(todoItem *TodoItem) error {
	return database.CreateItem(todoItem, repo.TableName)
}

func (repo *TodoItemRepo) GetTodoItems() (*[]TodoItem, error) {
	results, err := database.GetItems(repo.TableName)
	if err != nil {
		return nil, err
	}

	todoItems := make([]TodoItem, len(results))
	for i, item := range results {
		var todoItem TodoItem
		err = dynamodbattribute.UnmarshalMap(item, &todoItem)
		if err != nil {
			return nil, err
		}
		todoItems[i] = todoItem
	}

	return &todoItems, nil
}
