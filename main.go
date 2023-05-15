package main

import (
	"github.com/Xshen-aran/aran_platform/apps/router"
	_ "github.com/Xshen-aran/aran_platform/config"
)

func main() {
	router.GinRouter.Run(":8080")
}
