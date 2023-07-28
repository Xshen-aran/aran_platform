package middleware

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Xshen-aran/aran_platform/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const tokenExpireDuration = 7 * 24 * time.Hour

var customSecret = []byte(config.Env["JWT_SECRET"])

type CustomClaims struct {
	Username string `json:"user_name"`
	IsActive bool   `json:"is_active"`
	Auth     int    `json:"is_admin"`
	jwt.RegisteredClaims
}

func JWTMiddlewares() func(c *gin.Context) {
	return func(c *gin.Context) {
		r := c.Request.URL.Path
		pattern := "/api/v1/users"
		pattern_health := "health_check"
		matched, _ := regexp.MatchString(pattern, r)
		matched_health, _ := regexp.MatchString(pattern_health, r)
		if matched || matched_health {
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
			"message": "缺少鉴权字段",
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
	c.Set("auth", mc.Auth)
	c.Set("claims", mc)
	c.Next()
}

func GenToken(username string, isActive bool, Auth int) (string, error) {
	claims := CustomClaims{
		Username: username,
		IsActive: isActive,
		Auth:     Auth,
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
