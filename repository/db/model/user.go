package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int32 `gorm:"primaryKey;autoIncrement"`
	Name      string
	Password  string
	Email     string
	Phone     string
	School    string
	SchoolID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
