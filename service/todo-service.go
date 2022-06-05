package service

import (
	"fmt"
	"go1/dto"
	"go1/entity"
	"go1/repository"
	"log"

	"github.com/mashingan/smapping"
)

type TodoService interface {
	All(userID string) []entity.Todo
	FindByID(todoID int) entity.Todo
	Insert(t dto.TodoDTO) entity.Todo
	Update(todoID int, t dto.TodoDTO) entity.Todo
	Delete(t entity.Todo)
	IsAllowedToEdit(userID string, todoID int) bool
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (service *todoService) Insert(t dto.TodoDTO) entity.Todo {
	todo := entity.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.todoRepository.InsertTodo(todo)
	return res
}

func (service *todoService) Update(todoID int, t dto.TodoDTO) entity.Todo {
	todo := entity.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.todoRepository.UpdateTodo(todoID, todo)
	return res
}

func (service *todoService) Delete(t entity.Todo) {
	service.todoRepository.DeleteTodo(t)
}

func (service *todoService) All(userID string) []entity.Todo {
	return service.todoRepository.AllTodo(userID)
}

func (service *todoService) FindByID(todoID int) entity.Todo {
	return service.todoRepository.FindTodoByID(todoID)
}

func (service *todoService) IsAllowedToEdit(userID string, todoID int) bool {
	t := service.todoRepository.FindTodoByID(todoID)
	id := fmt.Sprintf("%v", t.UserID)
	return userID == id
}
