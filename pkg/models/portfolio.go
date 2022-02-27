package models

import (
	"time"

	"gorm.io/gorm"
)

type Portfolio struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"title"`
	Summary   string `gorm:"summary"`
	Url       string `gorm:"url"`
	ImageUrl  string `gorm:"image_url"`
	ImageAlt  string `gorm:"image_alt"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
