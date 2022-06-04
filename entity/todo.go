package entity

import (
	// "database/sql"
	// "encoding/json"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int            `gorm:"primary_key:auto_increment" json:"id"`
	Title     string         `gorm:"text" json:"title"`
	Isi       string         `gorm:"text" json:"isi"`
	Reminder  sql.NullTime   `gorm:"<-" json:"reminder,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime; <-:create" json:"createdAt"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime; <-:update" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleteAt"`
	ColorID   int            `gorm:"index" json:"colorId"`
	Color     Color          `gorm:"foreignkey:ColorID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"color"`
	UserID    int            `gorm:"not null" json:"userId"`
	User      User           `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}

func (Todo) TableName() string {
	return "todo"
}
