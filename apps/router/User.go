package router

import (
	"github.com/Xshen-aran/aran_platform/apps/controller"
	"github.com/Xshen-aran/aran_platform/apps/databases"
	"github.com/Xshen-aran/aran_platform/apps/modules"
	"github.com/gin-gonic/gin"
)

func userRouter(e *gin.Engine) {
	v := e.Group("/v1")
	{

		r := v.Group("/users")
		{

			r.POST("/all", controller.GetUsers)

		}
	}
}
func init() {
	include(userRouter)
	var count int64
	databases.Db.Debug().Model(&modules.RolePermissions{}).Count(&count)
	if count == 0 {
		// 初始化权限
		databases.Db.Debug().Create(&modules.RolePermissions{Role: "Tester", Permissions: modules.Tester})
		databases.Db.Debug().Create(&modules.RolePermissions{Role: "Developer", Permissions: modules.Developer})
		databases.Db.Debug().Create(&modules.RolePermissions{Role: "Manager", Permissions: modules.Manager})
		databases.Db.Debug().Create(&modules.RolePermissions{Role: "Admin", Permissions: modules.Admin})
	}
	// 创建管理员账户
	var admin modules.Users
	databases.Db.Debug().Where("username = ?", "admin").First(&admin)
	if admin.Id == 0 {
		admin = modules.Users{
			Username:          "admin",
			Password:          "Password@1",
			RolePermissionsId: modules.Admin,
		}
		admin.Creater()
	}
}
