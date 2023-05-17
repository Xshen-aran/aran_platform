package main

import (
	_ "github.com/Xshen-aran/aran_platform/apps/databases"
	"github.com/Xshen-aran/aran_platform/apps/router"
	_ "github.com/Xshen-aran/aran_platform/config"
)

func main() {
	router.ReloadRouter().Run(":8080")

}
