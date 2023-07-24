package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID		 uint    `json:"id" gorm:"primary_key;auto_increment;not null;comment:'主键 - 自增id'"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;type:datetime(0)"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;type:datetime(0);autoUpdateTime"`
	Name     string  `json:"name" gorm:"column:name;comment:'用户名'"`
	Email    string  `json:"email" gorm:"column:email;comment:'邮箱'"`
	Phone    int     `json:"phone" gorm:"column:phone;comment:'手机号'"`
	Sex      int     `json:"sex" gorm:"column:sex;comment:'性别'"`
	Token    string  `json:"token" gorm:"column:token;comment:'token'"`
	Password string  `json:"password" gorm:"column:password;comment:'密码'"`
	Wallet   float64 `json:"wallet" gorm:"column:wallet;comment:'钱包'"`
	Notify   int 	`json:"notify" gorm:"column:notify;comment:'是否开启通知'"`
	Type   	 int 	`json:"type" gorm:"column:type;comment:'类型 0 注册 1 登录 2 wx登录 3一键登录'"`
}

func UserDb(db *gorm.DB)  {
	db.AutoMigrate(User{})
}

