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
		logger.Logger.Sugar().Info("valid token", token)
		client_ip := ctx.ClientIP()
		logger.Logger.Sugar().Info(client_ip)
		claim, err := jwt.ParseToken(token)
		logger.Logger.Sugar().Info(claim)
		if err != nil {
			logger.Logger.Error("parse token err", zap.Error(err))
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.Forbidden,
				"errorMsg":  errors.GetMsg(errors.Forbidden) + " parse token err",
			})
			return
		}

		if claim.ExpiresAt < time.Now().Unix() {
			logger.Logger.Sugar().Warn("expires")
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.Forbidden,
				"errorMsg":  errors.GetMsg(errors.Forbidden) + " please login",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"errorCode": errors.SUCCESS,
			"errorMsg":  errors.GetMsg(errors.SUCCESS),
		})
	}
}
