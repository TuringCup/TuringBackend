package dao

import (
	"context"

	"github.com/TuringCup/TuringBackend/repository/db/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Create(&user).Error
}

func (dao *UserDao) ExistOrNotByUserName(name string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("name = ?", name).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("name = ?", name).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func (dao *UserDao) ExistOrNotByEmail(email string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}
