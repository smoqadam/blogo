package models

type Tag struct {
	ID  int `gorm:"primary_key"`
	Tag string
}
