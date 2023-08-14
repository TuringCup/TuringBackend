package service

import (
	"context"
	"errors"
	"strconv"

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

func FindUser(ctx context.Context, req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {

	userdao := dao.NewUserDao(ctx)
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		log.Error(err)
		return
	}
	user, err := userdao.FindUserById(id)
	if err != nil {
		log.Error(err)
	}
	resp = &types.GetUserResponse{
		ID:          int(user.ID),
		Name:        user.Name,
		Password:    user.Password,
		Phone:       user.Phone,
		Email:       user.Email,
		School:      user.School,
		SchoolId:    user.SchoolID,
		CreatedTime: user.CreatedAt.String(),
		UpdatedTime: user.UpdatedAt.String(),
	}
	return resp, err
}
