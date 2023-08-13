package service

import (
	"context"
	"errors"

	"github.com/TuringCup/TuringBackend/repository/db/dao"
	"github.com/TuringCup/TuringBackend/repository/db/model"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func UserReigster(ctx context.Context, req *types.RegisterRequest) (resp interface{}, err error) {
	userdao := dao.NewUserDao(ctx)

	// 检查用户名是否已经被注册
	_, exist, err := userdao.ExistOrNotByUserName(req.Username)
	if err != nil {
		log.Error(err)
		return
	}
	if exist {
		err = errors.New("用户名已经被注册")
		return
	}
	// 加密密码
	encrypt_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return
	}
	user := &model.User{
		Name:     req.Username,
		Password: string(encrypt_password),
		Email:    req.Email,
		School:   req.School,
		SchoolID: req.SchoolId,
		Phone:    req.Phone,
	}

	// 创建用户
	if err = userdao.CreateUser(user); err != nil {
		log.Error(err)
		return
	}
	return
}
