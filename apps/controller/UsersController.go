package controller

import (
	"github.com/Xshen-aran/aran_platform/apps/databases"
	"github.com/Xshen-aran/aran_platform/apps/modules"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var Users modules.Users
	databases.Db.Debug().First(&Users)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取用户列表成功",
		"data":    Users,
	})
}
