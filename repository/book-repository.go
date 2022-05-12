package repository

import (
	"belajar-API/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Book, error)
	FindById(Id int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(Id int, book entity.Book) (entity.Book, error)
	Delete(book entity.Book) (entity.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//---------------------------------------------------------
func (r *repository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindById(Id int) (entity.Book, error) {
	var book entity.Book
	err := r.db.Find(&book, Id).Error

	return book, err
}

func (r *repository) Create(book entity.Book) (entity.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(Id int, book entity.Book) (entity.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) Delete(book entity.Book) (entity.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
