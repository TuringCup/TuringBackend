package model

type CircleRaceInfo struct {
	Tid   int32 `gorm:"primaryKey"`
	Rid   int32 `gorm:"primaryKey"`
	Score int32
}
