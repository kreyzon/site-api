package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"username"`
	Password  string `gorm:"password"`
	Fullname  string `gorm:"fullname"`
	Role      string `gorm:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
