package controller

import (
	"github.com/Xshen-aran/aran_platform/apps/databases"
	"github.com/Xshen-aran/aran_platform/apps/middleware"
	"github.com/Xshen-aran/aran_platform/apps/modules"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	getUsers(c).send()
}
func getUsers(c *gin.Context) *response {
	var users []*modules.Users
	databases.Db.Debug().Find(&users)
	var user modules.Users
	databases.Db.Debug().Where(&user).FirstOrCreate(&user, map[string]interface{}{
		"username": user.Username,
		"password": user.Password,
	})

	for _, v := range users {
		v.Password = modules.SHA256(v.Password)
	}
	return &response{
		code: StatusOK,
		body: body{
			"code":    StatusOK,
			"message": statusText(StatusOK),
			"data":    users,
		},
		c: c,
	}
}

func Register(c *gin.Context) {
	register(c).send()
}
func register(c *gin.Context) *response {
	var users modules.Users
	if err := shouldBindJSON(c, &users); err != nil {
		return nil
	}
	if err := users.Creater(); err != nil {
		return &response{
			c:   c,
			err: err,
		}
	}
	token, _ := middleware.GenToken(users.Username, true, int(users.RolePermissionsId))
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	return &response{
		code: StatusOK,
		body: body{
			"code":    200,
			"message": "创建成功",
			"data":    users,
		},
		c: c,
	}
}

func Login(c *gin.Context) {
	login(c).send()
}

func login(c *gin.Context) *response {
	var login modules.Login
	if err := shouldBindJSON(c, &login); err != nil {
		return nil
	}

	return &response{
		code: StatusOK,
		body: body{
			"code":    200,
			"message": "获取用户列表成功",
			"data":    login,
		},
		c: c,
	}
}
