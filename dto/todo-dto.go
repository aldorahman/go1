package dto

import (
	// "time"

	// "encoding/json"
	"database/sql"
	// "time"
	// "gorm.io/datatypes"
)

// import "database/sql"

type TodoDTO struct {
	Title    string       `json:"title" form:"title" binding:"required"`
	Isi      string       `json:"isi" form:"title" binding:"required"`
	Reminder sql.NullTime `json:"reminder,omitempty" form:"reminder,omitempty"`
	ColorID  int          `json:"colorId,omitempty"  form:"colorId,omitempty"`
	UserID   int          `json:"userId,omitempty"  form:"userId,omitempty"`
}
