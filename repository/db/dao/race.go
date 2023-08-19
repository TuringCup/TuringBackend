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
func TestNewRaceDao() *RaceDao {
	return &RaceDao{TestDBClient()}
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

func (dao *RaceDao) FindRaceById(id int) (race *model.Race, err error) {
	err = dao.DB.Model(&model.Race{}).Where("id=?", id).First(&race).Error
	if err != nil {
		return nil, err
	}
	return race, nil
}

func (dao *RaceDao) GetRaceCount() (count int64, err error) {
	err = dao.DB.Model(&model.Race{}).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (dao *RaceDao) FindRaceByPage(page int32, perPage int32) (races []model.Race, err error) {
	if err = dao.DB.Model(&model.Race{}).Offset(int((page - 1) * perPage)).Limit(int(perPage)).Find(&races).Error; err != nil {
		return nil, err
	}
	return races, nil
}
