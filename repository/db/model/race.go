package model

import (
	"gorm.io/gorm"
	"time"
)

type Race struct {
	gorm.Model
	Id        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdTime"`
	UpdatedAt time.Time `json:"updatedTime"`
	DeletedAt gorm.DeletedAt
}
