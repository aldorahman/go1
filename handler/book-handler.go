package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"belajar-API/dto"
	"belajar-API/helper"
	"belajar-API/service"
)

type bookHandler struct {
	bookService service.Service
}

func NewBookHandler(bookService service.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "masuk",
		"data":   "OK",
	})
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []helper.BookResponse
	for _, b := range books {
		bookResponse := helper.ConvertBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	bookResponse := helper.ConvertBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	var bookRequest dto.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	book, err := h.bookService.Create(bookRequest)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest dto.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	b, err := h.bookService.Update(id, bookRequest)

	bookResponse := helper.ConvertBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	b, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	bookResponse := helper.ConvertBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data":   bookResponse,
		"status": "deleted",
	})
}
