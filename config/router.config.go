package config

import "fmt"

const (
	HomeApi = ""
	UserApi = ""
	LoginApi = ""
	RegisterApi = ""
)

const (
	HomeGroup = "/home"
	UserGroup = "/user"
	LoginGroup = "/login"
	RegisterGroup = "/register"
)


var AuthPath = []string{
	fmt.Sprintf("%s%s", UserGroup, UserApi),
	fmt.Sprintf("%s%s", LoginGroup, LoginApi),
	fmt.Sprintf("%s%s", RegisterGroup, RegisterApi),
}