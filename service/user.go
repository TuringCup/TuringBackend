package service

import (
	"context"
	"errors"
	"fmt"

	errs "github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/repository/db/dao"
	"github.com/TuringCup/TuringBackend/repository/db/model"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func UserReigster(ctx context.Context, req *types.RegisterRequest) (resp interface{}, err error) {
	userdao := dao.NewUserDao(ctx)

	// 检查用户名是否已经被注册
	_, exist, err := userdao.ExistOrNotByUserName(req.Username)
	if err != nil {
		fmt.Println(err)
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

func UserLogin(ctx context.Context, req *types.LoginRequest) (resp interface{}, err error) {
	userdao := dao.NewUserDao(ctx)
	user, exist, err := userdao.ExistOrNotByUserName(req.Username)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		return
	}
	if !exist {
		fmt.Fprintln(gin.DefaultWriter, req.Username+" not exist")
		err = errors.New(errs.GetMsg(errs.UserNotExist))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		err = errors.New(errs.GetMsg(errs.UserPasswordWrong))
		fmt.Fprintln(gin.DefaultWriter, req.Username+" password wrong")
		return
	}
	return
}
