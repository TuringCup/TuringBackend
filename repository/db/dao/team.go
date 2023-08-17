package dao

import (
	"context"
	"github.com/TuringCup/TuringBackend/repository/db/model"
	"gorm.io/gorm"
)

type TeamDao struct {
	*gorm.DB
}

func NewTeamDao(ctx context.Context) *TeamDao {
	return &TeamDao{NewDBClient(ctx)}
}

func TestNewTeamDao() *TeamDao {
	return &TeamDao{TestDBClient()}
}

func (dao *TeamDao) CreateTeam(team *model.Team) error {
	return dao.DB.Create(&team).Error
}

func (dao *TeamDao) ExistOrNotByTeamName(name string) (team *model.Team, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Race{}).Where("name=?", name).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.Race{}).Where("name=?", name).First(&team).Error
	if err != nil {
		return team, true, err
	}
	return team, true, nil
}
