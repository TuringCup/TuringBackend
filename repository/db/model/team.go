package model

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	gorm.Model
	ID        int32 `gorm:"primaryKey;autoIncrement"`
	Rid       int32
	Name      string
	CapId     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
