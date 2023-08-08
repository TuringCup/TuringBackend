package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	School    string    `json:"school"`
	SchoolID  string    `json:"schoolID"`
	CreatedAt time.Time `json:"createdTime"`
	UpdatedAt time.Time `json:"updatedTime"`
	DeletedAt gorm.DeletedAt
}
