package service

import (
	"belajar-API/dto"
	"belajar-API/entity"
	"belajar-API/repository"
)

type Service interface {
	FindAll() ([]entity.Book, error)
	FindById(Id int) (entity.Book, error)
	Create(bookrequest dto.BookRequest) (entity.Book, error)
	Update(Id int, bookrequest dto.BookRequest) (entity.Book, error)
	Delete(Id int) (entity.Book, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

//-------------------------------------------
func (s *service) FindAll() ([]entity.Book, error) {
	book, err := s.repository.FindAll()

	return book, err
}

func (s *service) FindById(Id int) (entity.Book, error) {
	book, err := s.repository.FindById(Id)

	return book, err
}

func (s *service) Create(bookrequest dto.BookRequest) (entity.Book, error) {
	price, err := bookrequest.Price.Int64()
	rating, err := bookrequest.Rating.Int64()
	discount, err := bookrequest.Discount.Int64()

	book := entity.Book{
		Title:       bookrequest.Title,
		Description: bookrequest.Description,
		Price:       int(price),
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}

func (s *service) Update(Id int, bookrequest dto.BookRequest) (entity.Book, error) {
	book, err := s.repository.FindById(Id)

	price, err := bookrequest.Price.Int64()
	rating, err := bookrequest.Rating.Int64()
	discount, err := bookrequest.Discount.Int64()

	book.Title = bookrequest.Title
	book.Description = bookrequest.Description
	book.Price = int(price)
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(Id, book)

	return newBook, err
}

func (s *service) Delete(Id int) (entity.Book, error) {
	book, err := s.repository.FindById(Id)

	Book, err := s.repository.Delete(book)

	return Book, err
}
