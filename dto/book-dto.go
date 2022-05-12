package dto

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"number"`
	Rating      json.Number `json:"rating" binding:"number"`
	Discount    json.Number `json:"discount" binding:"number"`
}
