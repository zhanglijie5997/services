package order

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	core "service/db"
	"service/db/order"
	"service/db/user"
	"service/models"
)

func Order(context *gin.Context) {
	_token := context.Request.Header["Token"]
	fmt.Print(context.Request.Header)

	var _user user.User
	// 查询用户
	core.DB.Where("token = ?", _token).First(&_user)
	if _token == nil {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpNeedToken])
		return
	}

	// 用户不存在
	if _user.ID == 0 {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpUserIsNotRegister])
		return
	}

	// 查询订单
	var _list []order.Order
	core.DB.Model(&order.Order{}).Limit(10).Where("send_id = @name OR  receive_id = @name", sql.Named("name", _user.ID)).Find(&_list)

	context.JSON(http.StatusOK, models.HttpModels{
		Message: config.HttpSuccessMessage,
		Data:    _list,
		Code:    config.HttpSuccess,
	})
}
