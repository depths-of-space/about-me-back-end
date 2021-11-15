package types

import (
	"time"

	"gorm.io/gorm"
)

// Tags public graph issues
type Tags struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	PostRefer uint
	Subject   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
