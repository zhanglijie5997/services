package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	core "service/db"
	"service/db/user"
	"service/models"
)

type Body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     int    `json:"type"`
	Code     string `json:"code"`
	Vertify  string `json:"vertify"`
}

func Login(context *gin.Context) {
	var _body Body
	var _user user.User
	err := context.BindJSON(&_body)
	if err != nil {
		return
	}
	_dbHaveUser := core.DB.Where("email = ?", _body.Email).First(&_user)
	fmt.Println(_dbHaveUser)

	context.JSON(http.StatusOK, models.HttpModels{
		Code:    config.HttpSuccess,
		Message: config.HttpSuccessMessage,
		Data:    _user,
	})
}
