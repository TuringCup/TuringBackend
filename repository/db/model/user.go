package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int32 `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string
	Phone     string
	School    string
	SchoolID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
