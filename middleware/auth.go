package middleware

import (
	"net/http"
	"time"

	"github.com/TuringCup/TuringBackend/pkg/errors"
	"github.com/TuringCup/TuringBackend/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		client_ip := ctx.ClientIP()
		claim, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"ErrorCode": errors.Forbidden,
				"ErrorMsg":  errors.GetMsg(errors.Forbidden),
			})
			ctx.Abort()
			return
		}
		if claim.IP != client_ip {
			ctx.JSON(http.StatusForbidden, gin.H{
				"ErrorCode": errors.Forbidden,
				"ErrorMsg":  errors.GetMsg(errors.Forbidden) + " ip changed",
			})
			ctx.Abort()
			return
		}
		if claim.ExpiresAt < time.Now().Unix() {
			ctx.JSON(http.StatusForbidden, gin.H{
				"ErrorCode": errors.TokenTimeout,
				"ErrorMsg":  errors.GetMsg(errors.TokenTimeout),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
