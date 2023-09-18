package model

import (
	"gorm.io/gorm"
	"time"
)

type Recipe struct {
	Id   uint `gorm:"primaryKey"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
