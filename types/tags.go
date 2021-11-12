package types

import (
	"time"

	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Subject   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
