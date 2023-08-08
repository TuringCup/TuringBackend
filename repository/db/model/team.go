package model

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	gorm.Model
	ID        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Rid       int32     `json:"rid"`
	Name      string    `json:"name"`
	CapId     int32     `json:"capId"`
	CreatedAt time.Time `json:"createdTime"`
	UpdatedAt time.Time `json:"updatedTime"`
}
