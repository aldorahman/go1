package dto

//ColorCreateDTO is is a model that clinet use when create a new book
type ColorDTO struct {
	ColorType string `json:"color_type" form:"color_type" binding:"required"`
	ColorName string `json:"color_name" form:"color_name" binding:"required"`
}
