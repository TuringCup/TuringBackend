package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/pkg/utils/jwt"
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
		if err != nil {
			ctx.JSON(http.StatusBadRequest, types.UploadFileResponse{
				ErrorMsg:  "上传失败",
				ErrorCode: errors.ERROR,
			})
			return
		}
		fmt.Println(file.Size)
		// 保证文件大小小于16MB
		if file.Size > (1 << 24) {
			ctx.JSON(http.StatusBadRequest, types.UploadFileResponse{
				ErrorMsg:  "上传失败,文件大小太大",
				ErrorCode: errors.ERROR,
			})
			return
		}
		fmt.Printf("%v\n", file.Filename)
		ctx.SaveUploadedFile(file, "./data/userfiles/"+file.Filename)
		ctx.JSON(http.StatusOK, types.UploadFileResponse{
			ErrorMsg:  "上传成功",
			ErrorCode: errors.SUCCESS,
		})
	}
}
