package middleware

import (
	"net/http"
	"time"

	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/pkg/utils/jwt"
	"github.com/TuringCup/TuringBackend/pkg/utils/logger"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token := ctx.Query("token")
		token := ctx.Request.FormValue("token")
		client_ip := ctx.ClientIP()
		claim, err := jwt.ParseToken(token)
		defer logger.Logger.Sync()
		logger.Logger.Sugar().Info(token)
		logger.Logger.Sugar().Error(claim)
		if err != nil {
			logger.Logger.Sugar().Error(err)
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.Forbidden,
				"errorMsg":  errors.GetMsg(errors.Forbidden),
			})
			ctx.Abort()
			return
		}
		if claim.IP != client_ip {
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.Forbidden,
				"errorMsg":  errors.GetMsg(errors.Forbidden) + " ip changed",
			})
			ctx.Abort()
			return
		}
		if claim.ExpiresAt < time.Now().Unix() {
			ctx.JSON(http.StatusForbidden, gin.H{
				"errorCode": errors.TokenTimeout,
				"errorMsg":  errors.GetMsg(errors.TokenTimeout),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
