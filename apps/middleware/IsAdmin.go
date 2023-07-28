package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdminAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		username := claims.(*CustomClaims)
		if username.Auth != 4 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "没有权限,请联系管理员",
				"data":    "",
			})
			c.Abort()
			return
		}

	}

}
