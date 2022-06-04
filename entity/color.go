package entity

//Color struct represents books table in database
type Color struct {
	ID        int    `gorm:"primary_key:auto_increment" json:"id"`
	ColorType string `gorm:"type:varchar(255)" json:"color_type"`
	ColorName string `gorm:"type:varchar(255)" json:"color_name"`
}

func (Color) TableName() string {
	return "color"
}
