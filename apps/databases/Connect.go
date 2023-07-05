package databases

import (
	"fmt"

	"github.com/Xshen-aran/aran_platform/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Env["DATABASE_USERNAME"], config.Env["DATABASE_PASSWORD"], config.Env["DATABASE_URL"], config.Env["DATABASE_PORT"], config.Env["DATABASE_DB_NAME"])
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect to database failed, error: ", err)
		fmt.Println("dsn: ", dsn)
		panic(err)
	}
}
