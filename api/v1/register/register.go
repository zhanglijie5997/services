package register

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	dbInit "service/db"
	"service/db/user"
	"service/models"
	"service/utils"
)

type Json struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     int    `json:"type"` // 0 注册 1 登录 2 wx登录 3一键登录
	Code     string `json:"code"`
	Vertify  string `json:"vertify"`
}

func Register(context *gin.Context) {
	var _body Json
	if err := context.BindJSON(&_body); err != nil {
		fmt.Println("err", err.Error())
		// 处理错误
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpTypeError])
		return
	}
	fmt.Println("types", _body.Type)
	// 邮箱错误
	if _body.Email == "" {
		context.JSON(
			http.StatusOK,
			config.HttpCodeMesaage[config.HttpNeedEmail],
		)
		return
	}

	if !utils.PhoneValidate(_body.Email) {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpEmailError])
		return
	}

	// 密码错误
	if _body.Password == "" {
		context.JSON(
			http.StatusOK,
			config.HttpCodeMesaage[config.HttpNeedPassword],
		)
		return
	}
	_user := user.User{
		Name:     _body.Email,
		Email:    _body.Email,
		Sex:      1,
		Phone:    0,
		Password: _body.Password,
		Token:    "123",
		Wallet:   0,
		Notify:   0,
	}
	_dbHaveUser := dbInit.DB.Where("email = ?", _body.Email).First(&_user)
	// 用户已存在
	if _dbHaveUser.RowsAffected > 0 {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpUserIsRegister])
		return
	}
	result := dbInit.DB.Create(&_user)
	if result.Error != nil {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpRegisterError])
		return
	}
	context.JSON(http.StatusOK, models.HttpModels{
		Code: config.HttpSuccess,
		Data: _user,
	})
}
