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
	Phone     string
	Email     string
	School    string
	SchoolID  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
