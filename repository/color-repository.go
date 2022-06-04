package repository

import (
	"go1/entity"

	"gorm.io/gorm"
)

type ColorRepository interface {
	InsertColor(c entity.Color) entity.Color
	UpdateColor(c entity.Color) entity.Color
	DeleteColor(c entity.Color)
	AllColor() []entity.Color
	FindColorByID(colorID int) entity.Color
}

type colorConnection struct {
	connection *gorm.DB
}

func NewColorRepository(dbConn *gorm.DB) ColorRepository {
	return &colorConnection{
		connection: dbConn,
	}
}

func (db *colorConnection) InsertColor(c entity.Color) entity.Color {
	db.connection.Save(&c)
	db.connection.Find(&c)
	return c
}

func (db *colorConnection) UpdateColor(c entity.Color) entity.Color {
	db.connection.Save(&c)
	db.connection.Find(&c)
	return c
}

func (db *colorConnection) DeleteColor(c entity.Color) {
	db.connection.Delete(&c)
}

func (db *colorConnection) FindColorByID(colorID int) entity.Color {
	var color entity.Color
	db.connection.Find(&color, colorID)
	return color
}

func (db *colorConnection) AllColor() []entity.Color {
	var colors []entity.Color
	db.connection.Find(&colors)
	return colors
}
