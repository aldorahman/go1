package service

import (
	"go1/dto"
	"go1/entity"
	"go1/repository"
	"log"

	"github.com/mashingan/smapping"
)

// todo service
type ColorService interface {
	Insert(c dto.ColorDTO) entity.Color
	Update(id int, c dto.ColorDTO) entity.Color
	Delete(c entity.Color)
	FindByID(colorID int) entity.Color
	All() []entity.Color
}

type colorService struct {
	colorRepository repository.ColorRepository
}

func NewColorService(colorRepo repository.ColorRepository) ColorService {
	return &colorService{
		colorRepository: colorRepo,
	}
}

func (service *colorService) Insert(c dto.ColorDTO) entity.Color {
	color := entity.Color{}
	err := smapping.FillStruct(&color, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.colorRepository.InsertColor(color)
	return res
}

func (service *colorService) Update(id int, c dto.ColorDTO) entity.Color {
	color := entity.Color{}
	err := smapping.FillStruct(&color, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	color.ID = id
	res := service.colorRepository.UpdateColor(color)
	return res
}

func (service *colorService) Delete(c entity.Color) {
	service.colorRepository.DeleteColor(c)
}

func (service *colorService) All() []entity.Color {
	return service.colorRepository.AllColor()
}

func (service *colorService) FindByID(colorID int) entity.Color {
	return service.colorRepository.FindColorByID(colorID)
}
