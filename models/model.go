package models

type BaseModel struct {
	ID uint `gorm:"primary_key" json:"id"`
}
