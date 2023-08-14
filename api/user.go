package api

import (
	"fmt"
	"net/http"

	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/service"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.RegisterRequest
		if err := ctx.Bind(&request); err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
			ctx.JSON(
				http.StatusOK,
				types.RegisterResponse{
					ErrorCode: errors.InvalidParams,
					ErrorMsg:  errors.GetMsg(errors.InvalidParams),
				},
			)
			return
		}
		response, err := service.UserReigster(ctx.Request.Context(), &request)
		if err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.LoginRequest
		if err := ctx.Bind(&req); err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
			resp := types.LoginResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		response, err := service.UserLogin(ctx, &req)
		if err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
			ctx.JSON(http.StatusOK, response)
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func UserRegisterValidCodeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ValidCodeRequest
		if err := ctx.Bind(&req); err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
			resp := types.ValidCodeResponse{
				ErrorCode: errors.ValidCodeError,
				ErrorMsg:  errors.GetMsg(errors.ValidCodeError),
			}
			ctx.JSON(http.StatusOK, resp)
		}
		response, err := service.UserReigsterSendValidCode(ctx, &req)
		if err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
		}
		ctx.JSON(http.StatusOK, response)
	}
}
