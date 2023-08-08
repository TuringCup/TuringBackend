package model

import "gorm.io/gorm"

type Round struct {
	gorm.Model
	Id         int32 `gorm:"primaryKey;autoIncrement"`
	Tid1       int32
	Tid2       int32
	TidWin     int32
	RoundType  int32
	RecordPath string
}
