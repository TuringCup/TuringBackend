package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 生产环境中的服务端通常都不会填 *，应当填写指定域名
			c.Header("Access-Control-Allow-Origin", origin)
			// 允许使用的HTTP METHOD
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			// 允许使用的请求头
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			// 允许客户端访问的响应头
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			// 是否需要携带认证信息 Credentials 可以是 cookies、authorization headers 或 TLS client certificates
			// 设置为true时，Access-Control-Allow-Origin不能为 *
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// 放行OPTION请求，但不执行后续方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 放行
		c.Next()
	}
}
