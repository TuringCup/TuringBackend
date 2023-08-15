package service

import (
	"context"
	"errors"

	"fmt"

	"strconv"

	"github.com/TuringCup/TuringBackend/pkg/email"
	errs "github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/pkg/utils/jwt"
	"github.com/TuringCup/TuringBackend/repository/cache"
	"github.com/TuringCup/TuringBackend/repository/db/dao"
	"github.com/TuringCup/TuringBackend/repository/db/model"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserReigsterSendValidCode(ctx context.Context, req *types.ValidCodeRequest) (resp interface{}, err error) {
	code, err := cache.GenerateValidCode()
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.ValidCodeResponse{
			ErrorCode: errs.ValidCodeGenError,
			ErrorMsg:  errs.GetMsg(errs.ValidCodeGenError),
		}
		return
	}
	fmt.Fprintln(gin.DefaultWriter, req.Email)
	err = email.SendValidCode(req.Email, code)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.ValidCodeResponse{
			ErrorCode: errs.SendValidCodeError,
			ErrorMsg:  errs.GetMsg(errs.SendValidCodeError),
		}
		return
	}

	resp = types.ValidCodeResponse{
		ErrorCode: errs.SUCCESS,
		ErrorMsg:  errs.GetMsg(errs.SUCCESS),
	}
	return
}

func UserReigster(ctx context.Context, req *types.RegisterRequest) (resp interface{}, err error) {
	userdao := dao.NewUserDao(ctx)

	// 检查用户名是否已经被注册
	_, exist, err := userdao.ExistOrNotByUserName(req.Username)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.RegisterResponse{
			ErrorCode: errs.RegisterFailed,
			ErrorMsg:  errs.GetMsg(errs.RegisterFailed),
		}
		return
	}
	if exist {
		err = errors.New("用户名已经被注册")
		resp = types.RegisterResponse{
			ErrorCode: errs.UserNameUsed,
			ErrorMsg:  errs.GetMsg(errs.UserNameUsed),
		}
		return
	}
	_, exist, err = userdao.ExistOrNotByEmail(req.Email)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.RegisterResponse{
			ErrorCode: errs.RegisterFailed,
			ErrorMsg:  errs.GetMsg(errs.RegisterFailed),
		}
		return
	}

	if exist {
		err = errors.New("邮箱已经被注册")
		resp = types.RegisterResponse{
			ErrorCode: errs.EmailUsed,
			ErrorMsg:  errs.GetMsg(errs.EmailUsed),
		}
		return
	}

	// 校验邮箱验证码
	err = cache.CheckValidCode(req.ValidCode)

	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.RegisterResponse{
			ErrorCode: errs.ValidCodeError,
			ErrorMsg:  errs.GetMsg(errs.ValidCodeError),
		}
		return
	}

	// 加密密码
	encrypt_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.RegisterResponse{
			ErrorCode: errs.RegisterFailed,
			ErrorMsg:  errs.GetMsg(errs.RegisterFailed),
		}
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
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.RegisterResponse{
			ErrorCode: errs.RegisterFailed,
			ErrorMsg:  errs.GetMsg(errs.RegisterFailed),
		}
		return
	}

	resp = types.RegisterResponse{
		ErrorCode: errs.SUCCESS,
		ErrorMsg:  errs.GetMsg(errs.SUCCESS),
	}
	return
}

func UserLogin(ctx *gin.Context, req *types.LoginRequest) (resp interface{}, err error) {

	userdao := dao.NewUserDao(ctx)
	user, exist, err := userdao.ExistOrNotByUserName(req.Username)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		resp = types.LoginResponse{
			ErrorCode: errs.LoginFailed,
			ErrorMsg:  errs.GetMsg(errs.LoginFailed),
		}
		return
	}
	if !exist {
		fmt.Fprintln(gin.DefaultWriter, req.Username+" not exist")
		err = errors.New(errs.GetMsg(errs.UserNotExist))
		resp = types.LoginResponse{
			ErrorCode: errs.UserNotExist,
			ErrorMsg:  errs.GetMsg(errs.UserNotExist),
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, req.Username+" password wrong"+err.Error())
		resp = types.LoginResponse{
			ErrorCode: errs.UserPasswordWrong,
			ErrorMsg:  errs.GetMsg(errs.UserPasswordWrong),
		}
		return
	}
	access_token, refresh_token, err := jwt.GenerateToken(int(user.ID), user.Name, ctx.ClientIP())
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, req.Username+" password wrong"+err.Error())
		resp = types.LoginResponse{
			ErrorCode: errs.GenerateTokenError,
			ErrorMsg:  errs.GetMsg(errs.GenerateTokenError),
		}
		return
	}
	resp = types.LoginResponse{
		Token:        access_token,
		RefreshToken: refresh_token,
		ErrorCode:    errs.SUCCESS,
		ErrorMsg:     errs.GetMsg(errs.SUCCESS),
	}
	return
}

func FindUser(ctx context.Context, req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {

	userdao := dao.NewUserDao(ctx)
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return resp, err
	}
	user, err := userdao.FindUserById(id)
	if err != nil {
		return resp, err
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
