package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	todoitem "todo-api/app/entity"
)

var once sync.Once
var Db *gorm.DB

func ConfigurePostgresDatabase() {
	once.Do(func() {
		dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Error to connect to database")
		}
		Db = db
	})

	err := Db.AutoMigrate(&todoitem.TodoItemEntity{})
	if err != nil {
		panic("Error to migrate database")
	}
}
