package models

import "gorm.io/gorm"

type Book struct {
	*gorm.Model
	Name string
	AuthorID int
	Author Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}