package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"service/config"
	core "service/db"
	"service/db/order"
	"service/db/user"
	"service/models"
)

type CreateOrderBody struct {
	Id          string `json:"id"`
	PayType     int    `json:"payType"`
	ProductCode string `json:"productCode"`
	Desc        string `json:"desc"`
	Code        string `json:"code"`
	Type        int    `json:"type"`
	OrderNum    int    `json:"orderNum"`
	CouponId    int32  `json:"couponId"`
}

func CreateOrder(context *gin.Context) {
	var _body CreateOrderBody
	_token := context.Request.Header.Get("token")
	err := context.BindJSON(&_body)
	if err != nil {
		return
	}
	// 先查询订单
	var _sendOrder order.Order
	core.DB.Model(&order.Order{}).Where("id = ?", _body.Id).Last(&_sendOrder)
	// 订单不存在
	if _sendOrder.OrderNo == "" {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpOrderNotFound])
		return
	}
	_uuid, _ := uuid.NewUUID()
	// 根据token 查用户信息
	var _user user.User
	core.DB.Model(&user.User{}).Where("token = ?", _token).First(&_user)
	// 购买用户不存在
	if _user.Token != _token || _token == "" {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpUserIsNotRegister])
		return
	}
	// 根据发布订单id查询发布商品用户
	var _receiveUser user.User
	core.DB.Preload("orders").Preload("orders").Where("id = ?", _body.Id).Last(&_receiveUser)
	fmt.Println(_receiveUser, "用户信息")

	// TODO:根据id查询订单物品信息
	// 订单类型名称
	var _orderType string = ""
	if _body.Type == 1 {
		_orderType = ""
	} else {
		_orderType = ""
	}

	_order := order.Order{
		Type:          1,
		OrderNo:       _uuid.String(),
		OrderName:     _orderType,
		OrderNumber:   _body.OrderNum,
		OrderCouponId: _body.CouponId,
		OrderStatus:   2,
	}
	fmt.Println(_order)
	//core.DB.Create(&_order)
	context.JSON(http.StatusOK, models.HttpModels{
		Code: config.HttpSuccess,
		Data: _user,
	})
}
