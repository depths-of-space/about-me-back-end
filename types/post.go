package types

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Title     string
	Body      string
	Tags      []Tags
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
