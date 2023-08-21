package v1

import (
	"github.com/gin-gonic/gin"
	"service/api/v1/home"
	"service/api/v1/login"
	"service/api/v1/order"
	"service/api/v1/product"
	"service/api/v1/register"
	"service/api/v1/user"
	"service/config"
)

func RouterGroupV1(group *gin.Engine) {
	v1 := group.Group("v1")
	{
		_home := v1.Group(config.HomeGroup)
		{
			_home.GET(config.HomeApi, home.Home)
		}
		_user := v1.Group(config.UserGroup)
		{
			_user.POST(config.UserApi, user.User)
		}
		_login := v1.Group(config.LoginGroup)
		{
			_login.POST(config.LoginApi, login.Login)
		}
		_register := v1.Group(config.RegisterGroup)
		{
			_register.POST(config.RegisterApi, register.Register)
		}
		_order := v1.Group(config.OrderGroup)
		{
			_order.GET(config.OrderApi, order.Order)
			_order.POST(config.CreateOrderApi, order.CreateOrder)
		}
		_product := v1.Group(config.ProductGroup)
		{
			_product.GET(config.ProductDetailApi, product.Detail)
		}
	}
}
