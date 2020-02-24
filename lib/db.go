package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Model struct {
	DB *gorm.DB
}

func NewModel() *Model {
	m := Model{}
	// TODO: Make it configurable
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	m.DB = db
	return &m
}
