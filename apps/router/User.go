package router

import (
	"github.com/gin-gonic/gin"
)

func userRouter(e *gin.Engine) {
	r := e.Group("/users")
	{

		r.POST("/register", func(ctx *gin.Context) {
			var data map[string]interface{}
			ctx.BindJSON(&data)
			ctx.JSON(200, gin.H{
				"message": "test",
				"data":    data,
			})
		})
		r.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "test",
			})
		})
		r.GET("/adminlogin", func(ctx *gin.Context) {
			ctx.SetCookie("token", "test", 3600, "/", "localhost", false, true)
			ctx.JSON(200, gin.H{
				"message": "test",
			})
		})
	}
}
func init() {
	include(userRouter)
}
