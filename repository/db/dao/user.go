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
func TestNewUserDao() *UserDao {
	return &UserDao{TestDBClient()}
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

func (dao *UserDao) ExistOrNotBySchoolID(id string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("school_id = ?", id).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("school_id = ?", id).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func (dao *UserDao) FindUserById(id int) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao *UserDao) UpdateUser(id int, user *model.User) (bool, error) {
	userInDb, err := dao.FindUserById(id)
	if err != nil {
		return false, err
	}
	err = dao.DB.Model(&model.User{}).Where("id=?", userInDb.ID).Updates(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
