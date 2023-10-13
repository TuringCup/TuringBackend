package api

import (
	"net/http"
	"time"

	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/pkg/utils/jwt"
	"github.com/TuringCup/TuringBackend/pkg/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ValidTokenHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// response := types.ValidTokenResponse{
		// 	ErrorCode: errors.SUCCESS,
		// 	ErrorMsg:  errors.GetMsg(errors.SUCCESS),
		// }
		// ctx.JSON(http.StatusOK, response)
		defer logger.Logger.Sync()
		token := ctx.Request.FormValue("token")
		logger.Logger.Info("valid token", zap.Any("token", token))
		client_ip := ctx.ClientIP()
		claim, err := jwt.ParseToken(token)
		if err != nil {
			logger.Logger.Error("parse token err", zap.Error(err))
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.Forbidden,
				"errorMsg":  errors.GetMsg(errors.Forbidden) + " parse token err",
			})
			return
		}
		if claim.IP != client_ip {
			logger.Logger.Error("parse token err", zap.Error(err))
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.Forbidden,
				"errorMsg":  errors.GetMsg(errors.Forbidden) + " ip changed",
			})
			return
		}
		if claim.ExpiresAt < time.Now().Unix() {
			logger.Logger.Error("parse token err", zap.Error(err))
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.Forbidden,
				"errorMsg":  errors.GetMsg(errors.Forbidden) + " please login",
			})
			return
		}
		logger.Logger.Info("parse token err", zap.Error(err))
		ctx.JSON(http.StatusForbidden, gin.H{
			"errorCode": errors.SUCCESS,
			"errorMsg":  errors.GetMsg(errors.SUCCESS),
		})
	}
}
