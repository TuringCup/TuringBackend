package model

import (
	"time"

	"gorm.io/gorm"
)

type Race struct {
	gorm.Model
	ID        int32 `gorm:"primaryKey;autoIncrement"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
