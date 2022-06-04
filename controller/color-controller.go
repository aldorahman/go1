package controller

import (
	"go1/dto"
	"go1/entity"
	"go1/helper"
	"go1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ColorController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type colorController struct {
	colorService service.ColorService
	jwtService   service.JWTService
}

func NewColorController(colorServ service.ColorService, jwtServ service.JWTService) ColorController {
	return &colorController{
		colorService: colorServ,
		jwtService:   jwtServ,
	}
}

func (c *colorController) All(context *gin.Context) {
	var colors []entity.Color = c.colorService.All()
	res := helper.BuildResponse(true, "OK", colors)
	context.JSON(http.StatusOK, res)
}

func (c *colorController) FindByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var color entity.Color = c.colorService.FindByID(id)
	if color.ID == id {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", color)
		context.JSON(http.StatusOK, res)
	}
}

func (c *colorController) Insert(context *gin.Context) {
	var colorDTO dto.ColorDTO
	errDTO := context.ShouldBind(&colorDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		_, errToken := c.jwtService.ValidateToken(authHeader)
		if errToken != nil {
			panic(errToken.Error())
		}

		result := c.colorService.Insert(colorDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *colorController) Update(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var colorDTO dto.ColorDTO
	errDTO := context.ShouldBind(&colorDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.colorService.Update(id, colorDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *colorController) Delete(context *gin.Context) {
	var color entity.Color
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	color.ID = id
	c.colorService.Delete(color)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
