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

func (dao *TeamDao) GetTeamCount() (count int64, err error) {
	err = dao.DB.Model(&model.Team{}).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (dao *TeamDao) FindTeamByPage(page int32, perPage int32) (teams []model.Team, err error) {
	if err = dao.DB.Model(&model.Team{}).Offset(int((page - 1) * perPage)).Limit(int(perPage)).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (dao *TeamDao) FindTeamById(id int32) (team *model.Team, err error) {
	err = dao.DB.Model(&model.Team{}).Where("id=?", id).First(&team).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (dao *TeamDao) FindTeamByRidAndTid(rid int32, tid int32) (team *model.Team, err error) {
	err = dao.DB.Model(&model.CircleRaceInfo{}).Where("rid=? AND tid=?", rid, tid).Find(&team).Error
	if err != nil {
		return nil, err
	}
	return team, nil
}
