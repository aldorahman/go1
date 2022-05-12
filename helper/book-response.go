package helper

import "belajar-API/entity"

type BookResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
}

func ConvertBookResponse(b entity.Book) BookResponse {
	return BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
