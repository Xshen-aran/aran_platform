package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const tokenExpireDuration = 7 * 24 * time.Hour

var customSecret = []byte("aran_platform")

type CustomClaims struct {
	Username string `json:"user_name"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func JWTMiddlewares() func(c *gin.Context) {
	return func(c *gin.Context) {
		if r := c.Request.URL.Path; r == "/users/register/" || r == "/users/login/" || r == "/users/adminlogin/" {
			c.Next()
		} else {
			jwtHandler(c)
		}
	}
}

func jwtHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "请求头中没有Authorization字段",
		})
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Authorization字段 有误",
		})
		c.Abort()
		return
	}
	mc, err := parseToken(parts[1])
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效的token",
			},
		)
		c.Abort()
		return
	}
	c.Set("username", mc.Username)
	c.Set("claims", mc)
	c.Next()
}

func GenToken(username string, isActive, isAdmin bool) (string, error) {
	claims := CustomClaims{
		Username: username,
		IsActive: isActive,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
			Issuer:    "aran",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(customSecret)
}

func parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return customSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
