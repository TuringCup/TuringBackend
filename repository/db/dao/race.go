package dao

import (
	"context"
	"github.com/TuringCup/TuringBackend/repository/db/model"
	"gorm.io/gorm"
)

type RaceDao struct {
	*gorm.DB
}

func NewRaceDao(ctx context.Context) *RaceDao {
	return &RaceDao{NewDBClient(ctx)}
}

func (dao *RaceDao) CreateRace(race *model.Race) error {
	return dao.DB.Create(&race).Error
}

func (dao *RaceDao) ExistOrNotByRaceName(name string) (race *model.Race, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Race{}).Where("name=?", name).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.Race{}).Where("name=?", name).First(&race).Error
	if err != nil {
		return race, true, err
	}
	return race, true, nil
}
