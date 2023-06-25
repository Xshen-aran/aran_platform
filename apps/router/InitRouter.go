package router

import (
	"fmt"

	"github.com/Xshen-aran/aran_platform/apps/middleware"
	"github.com/gin-gonic/gin"
)

type Function func(*gin.Engine)

var (
	options   = []Function{}
	GinRouter *gin.Engine
)

func init() {
	if err := middleware.InitLogger(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		panic(err)
	}
	GinRouter = gin.Default()
	GinRouter.Use(
		middleware.JWTMiddlewares(),
		middleware.GinLogger(),
		middleware.GinRecovery(true),
	)
}
func include(opts ...Function) {
	options = append(options, opts...)
}
func ReloadRouter() *gin.Engine {
	for _, v := range options {
		v(GinRouter)
	}
	return GinRouter
}
