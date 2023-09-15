package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	*gorm.Model
	Name       string
	FamilyName string
	Rank       int
	BirtDate   time.Time
}
