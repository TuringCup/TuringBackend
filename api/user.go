package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/pkg/utils/jwt"
	"github.com/TuringCup/TuringBackend/pkg/utils/logger"
	"github.com/TuringCup/TuringBackend/service"
	"github.com/TuringCup/TuringBackend/types"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.RegisterRequest
		defer logger.Logger.Sync()
		if err := ctx.Bind(&request); err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
			logger.Logger.Sugar().Error(err)
			ctx.JSON(
				http.StatusOK,
				types.RegisterResponse{
					ErrorCode: errors.InvalidParams,
					ErrorMsg:  errors.GetMsg(errors.InvalidParams),
				},
			)
			return
		}
		logger.Logger.Sugar().Info(request)
		response, err := service.UserReigster(ctx.Request.Context(), &request)
		logger.Logger.Sugar().Info(response)
		if err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.LoginRequest
		defer logger.Logger.Sync()
		if err := ctx.Bind(&req); err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
			logger.Logger.Sugar().Error(err)
			resp := types.LoginResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		logger.Logger.Sugar().Info(req)
		response, err := service.UserLogin(ctx, &req)
		logger.Logger.Sugar().Info(response)
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
		defer logger.Logger.Sync()
		if err := ctx.Bind(&req); err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
			logger.Logger.Sugar().Error(err)
			resp := types.ValidCodeResponse{
				ErrorCode: errors.ValidCodeError,
				ErrorMsg:  errors.GetMsg(errors.ValidCodeError),
			}
			ctx.JSON(http.StatusOK, resp)
		}
		logger.Logger.Sugar().Info(req)
		response, err := service.UserReigsterSendValidCode(ctx, &req)
		logger.Logger.Sugar().Info(response)
		if err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, err)
		}
		ctx.JSON(http.StatusOK, response)
	}
}
func UserFindHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.GetUserRequest
		request.ID = ctx.Param("id")

		token := ctx.Query("token")
		claim, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusOK, types.GetUserResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			})
			return
		}
		if id, err := strconv.Atoi(request.ID); claim.ID != id || err != nil {
			ctx.JSON(http.StatusOK, types.GetUserResponse{
				ErrorCode: errors.Forbidden,
				ErrorMsg:  errors.GetMsg(errors.Forbidden),
			})
			return
		}
		user, err := service.UserFind(ctx.Request.Context(), &request)
		if err != nil {
			ctx.JSON(http.StatusOK, types.GetUserResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			})
			return
		}
		user.ErrorCode = errors.SUCCESS
		user.ErrorMsg = errors.GetMsg(errors.SUCCESS)
		ctx.JSON(http.StatusOK, user)
	}
}

func UserUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateUserRequest
		urlId := ctx.Param("id")
		token := ctx.Query("token")
		claim, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusOK, types.GetUserResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			})
			return
		}
		if id, err := strconv.Atoi(urlId); claim.ID != id || err != nil {
			ctx.JSON(http.StatusOK, types.GetUserResponse{
				ErrorCode: errors.Forbidden,
				ErrorMsg:  errors.GetMsg(errors.Forbidden),
			})
			return
		}
		if err := ctx.Bind(&req); err != nil {
			resp := types.UpdateUserResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		id, err := strconv.Atoi(urlId)
		if err != nil {
			resp := types.UpdateUserResponse{
				ErrorCode: errors.InvalidParams,
				ErrorMsg:  errors.GetMsg(errors.InvalidParams),
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		req.ID = id
		resp, _ := service.UpdateUser(ctx, &req)
		ctx.JSON(http.StatusOK, *resp)
	}
}

func UserUploadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		logger.Logger.Sugar().Info(file)
		defer logger.Logger.Sync()
		if err != nil {
			logger.Logger.Sugar().Error(err)
			ctx.JSON(http.StatusBadRequest, types.UploadFileResponse{
				ErrorMsg:  "上传失败",
				ErrorCode: errors.ERROR,
			})
			return
		}
		logger.Logger.Sugar().Info(file.Size)
		// 保证文件大小小于16MB
		if file.Size > (1 << 24) {
			ctx.JSON(http.StatusBadRequest, types.UploadFileResponse{
				ErrorMsg:  "上传失败,文件大小太大",
				ErrorCode: errors.ERROR,
			})
			return
		}
		if filepath.Ext(file.Filename) != ".zip" {
			ctx.JSON(http.StatusBadRequest, types.UploadFileResponse{
				ErrorMsg:  "上传失败,文件格式错误",
				ErrorCode: errors.ERROR,
			})
			return
		}
		token := ctx.Query("token")
		if token == "" {
			token = ctx.Request.FormValue("token")
		}
		claim, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, types.UploadFileResponse{
				ErrorMsg:  "上传失败",
				ErrorCode: errors.ERROR,
			})
			return
		}
		filename := claim.Username + "_" + strconv.Itoa(claim.ID) + filepath.Ext(file.Filename)

		fmt.Printf("%v\n", file.Filename)
		fmt.Printf("%v\n", claim)
		fmt.Println(filename)
		ctx.SaveUploadedFile(file, "./data/userfiles/"+filename)
		logger.Logger.Sugar().Info(filename, " saved")
		ctx.JSON(http.StatusOK, types.UploadFileResponse{
			ErrorMsg:  "上传成功",
			ErrorCode: errors.SUCCESS,
		})
	}
}
