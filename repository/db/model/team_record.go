package model

import "gorm.io/gorm"

type TeamRecord struct {
	gorm.Model
	Id     int32 `gorm:"primaryKey;autoIncrement"`
	RaceId int32
	Uid    int32
	Tid    int32
}
