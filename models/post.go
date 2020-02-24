package models

import "time"

type Post struct {
	ID        int       `gorm:"primary_key"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tags      []Tag     `json:"tags" gorm:"many2many:posts_tags";"foreign_key:ID"`
}
