package repository

import (
	"go1/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	AllTodo(userID string) []entity.Todo
	FindTodoByID(todoID int) entity.Todo
	DeleteTodo(t entity.Todo)
	InsertTodo(t entity.Todo) entity.Todo
	UpdateTodo(t entity.Todo) entity.Todo
}

type todoConnection struct {
	connection *gorm.DB
}

func NewTodoRepository(dbConn *gorm.DB) TodoRepository {
	return &todoConnection{
		connection: dbConn,
	}
}

func (db *todoConnection) InsertTodo(t entity.Todo) entity.Todo {
	db.connection.Save(&t)
	db.connection.Preload("User").Preload("Color").Find(&t)
	return t
}

func (db *todoConnection) UpdateTodo(t entity.Todo) entity.Todo {
	db.connection.Debug().Save(&t)

	db.connection.Preload("User").Preload("Color").Find(&t, t.ID)
	return t
}

func (db *todoConnection) DeleteTodo(t entity.Todo) {
	db.connection.Delete(&t)
}

func (db *todoConnection) FindTodoByID(todoID int) entity.Todo {
	var todo entity.Todo
	db.connection.Preload("User").Preload("Color").Find(&todo, todoID)
	return todo
}

func (db *todoConnection) AllTodo(userID string) []entity.Todo {
	var todos []entity.Todo
	db.connection.Preload("User").Preload("Color").Where("user_id = ?", userID).Find(&todos)
	return todos
}
