package controller

import (
	"fmt"
	"go1/dto"
	"go1/entity"
	"go1/helper"
	"go1/service"
	"net/http"
	"strconv"

	// "time"

	// "strings"
	// "time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TodoController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type todoController struct {
	todoService service.TodoService
	jwtService  service.JWTService
}

func NewTodoController(todoServ service.TodoService, jwtServ service.JWTService) TodoController {
	return &todoController{
		todoService: todoServ,
		jwtService:  jwtServ,
	}
}

func (c *todoController) All(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	var todos []entity.Todo = c.todoService.All(userID)
	res := helper.BuildResponse(true, "OK", todos)
	context.JSON(http.StatusOK, res)
	return
}

func (c *todoController) FindByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	var todo entity.Todo = c.todoService.FindByID(id)
	if (todo == entity.Todo{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", todo)
		context.JSON(http.StatusOK, res)
	}
}

func (c *todoController) Insert(context *gin.Context) {
	var todoCreateDTO dto.TodoDTO
	errDTO := context.ShouldBind(&todoCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.Atoi(userID)
		if err == nil {
			todoCreateDTO.UserID = convertedUserID
		}
		result := c.todoService.Insert(todoCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *todoController) Update(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	var todoUpdateDTO dto.TodoDTO
	errDTO := context.ShouldBind(&todoUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.todoService.IsAllowedToEdit(userID, id) {
		user_id, errID := strconv.Atoi(userID)
		if errID == nil {
			todoUpdateDTO.UserID = user_id
		}
		result := c.todoService.Update(id, todoUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *todoController) Delete(context *gin.Context) {
	var todo entity.Todo
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	todo.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.todoService.IsAllowedToEdit(userID, todo.ID) {
		c.todoService.Delete(todo)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *todoController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
