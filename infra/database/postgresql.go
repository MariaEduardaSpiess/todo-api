package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"todo-api/app/entity"
	"todo-api/config"
)

var once sync.Once
var Db *gorm.DB

func InitDatabase() {
	once.Do(func() {
		env := config.Env.Database
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", env.Host, env.User, env.Pass, env.Name, env.Port, env.SslMode, env.TimeZone)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Error to connect to database")
		}
		Db = db
	})

	err := Db.AutoMigrate(&entity.TodoItem{})
	if err != nil {
		panic("Error to migrate database")
	}
}
