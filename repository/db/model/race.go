package model

import (
	"gorm.io/gorm"
	"time"
)

type Race struct {
	gorm.Model
	ID        int32 `gorm:"primaryKey;autoIncrement"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
