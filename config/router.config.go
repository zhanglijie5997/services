package config

import "fmt"

const (
	HomeApi          = ""
	UserApi          = ""
	LoginApi         = ""
	RegisterApi      = ""
	OrderApi         = ""
	ProductDetailApi = "/detail"
)

const (
	CreateOrderApi = "/create"
)

const (
	HomeGroup     = "/home"
	UserGroup     = "/user"
	LoginGroup    = "/login"
	RegisterGroup = "/register"
	OrderGroup    = "/order"
	ProductGroup  = "/product"
)

var AuthPath = []string{
	fmt.Sprintf("%s%s", UserGroup, UserApi),
	fmt.Sprintf("%s%s", LoginGroup, LoginApi),
	fmt.Sprintf("%s%s", RegisterGroup, RegisterApi),
}
