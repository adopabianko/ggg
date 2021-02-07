package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UUID struct {
	Id uuid.UUID `gorm:"primary_key;" json:"id"`
}
type Book struct {
	UUID
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Isbn        string `json:"isbn"`
}

func (b *Book) TableName() string {
	return "book"
}

func (u *UUID) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()

	return scope.SetColumn("id", uuid)
}
