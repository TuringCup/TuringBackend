package api

import (
	"fmt"
	"net/http"

	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/service"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.RegisterRequest
		if err := ctx.Bind(&request); err != nil {
			ctx.String(errors.InvalidParams, errors.GetMsg(errors.InvalidParams))
			ctx.JSON(
				http.StatusOK,
				types.RegisterResponse{
					ErrorCode: errors.InvalidParams,
					ErrorMsg:  errors.GetMsg(errors.InvalidParams),
				},
			)
		}
		response, err := service.UserReigster(ctx.Request.Context(), &request)
		if err != nil {
			log.Error(err)
			ctx.JSON(
				http.StatusOK,
				response,
			)
			return
		}
		ctx.JSON(http.StatusOK, types.RegisterResponse{
			ErrorCode: errors.SUCCESS,
			ErrorMsg:  errors.GetMsg(errors.SUCCESS),
		})
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.LoginRequest
		if err := ctx.Bind(&req); err != nil {
			fmt.Println(err)
			resp := types.LoginResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
		}

	}
}

func UserRegisterValidCodeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ValidCodeRequest
		if err := ctx.Bind(req); err != nil {
			resp := types.ValidCodeResponse{
				ErrorCode: errors.ValidCodeError,
				ErrorMsg:  errors.GetMsg(errors.ValidCodeError),
			}
			ctx.JSON(http.StatusOK, resp)
		}

	}
}
