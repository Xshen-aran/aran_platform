package modules

import (
	"time"

	"github.com/Xshen-aran/aran_platform/apps/databases"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permissions int

const (
	Tester Permissions = iota + 1
	Developer
	Manager
	Admin
)

type RolePermissions struct {
	Id          uint        `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Role        string      `gorm:"type:varchar(18);not null" json:"role"`
	Permissions Permissions `json:"permissions" gorm:"index"`
}

func (*RolePermissions) TableName() string {
	return "aran_role_permissions"
}

type Users struct {
	Id                uint            `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Uuid              uuid.UUID       `gorm:"type:varchar(36);uniqueIndex" json:"uuid"`
	Username          string          `gorm:"type:varchar(18);not null" json:"user_name" binding:"required,min=6,max=18"`
	Password          string          `gorm:"type:varchar(30);not null" json:"password" binding:"required,min=6,max=30"`
	LastLoginTime     time.Time       `gorm:"default:NULL" binding:"-" json:"last_login_time"`
	LoginOutTime      time.Time       `gorm:"default:NULL" binding:"-" json:"login_out_time"`
	RolePermissions   RolePermissions `gorm:"references:Permissions" json:"-"`
	RolePermissionsId Permissions     `json:"permissions" binding:"required"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"delete_at"`
}

func (*Users) TableName() string {
	return "aran_users"
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.Uuid = uuid.New()
	return
}

func (u *Users) Creater() error {
	return databases.Db.Debug().Create(u).Error
}

func (u *Users) Updater() error {
	return databases.Db.Debug().Updates(u).Error
}

func (u *Users) Deleter() error {
	return databases.Db.Debug().Delete(u).Error
}

type Login struct {
	Users           `gorm:"embedded" json:"users"`
	ConfirmPassword string `gorm:"-" json:"confirm_password" binding:"required"`
	Token           string
}

func init() {
	databases.Migrate(&Users{}, &RolePermissions{})
}
