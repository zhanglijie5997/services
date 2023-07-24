package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	"service/models"
	"strings"
)

// isHave 判断是否需要登陆
//
// path 路径
//
// @returns true -> 需要登陆 false -> 不需要登陆
func isHave(path string) bool {
	for _, item := range config.AuthPath {
		if strings.Contains(path, item) {
			return true
		}
	}
	return false
}


func Middlewares(next *gin.Context)  {
	if next.Request.Method == http.MethodGet {
		next.Next()
		return
	}
	_token := next.Request.Header.Get("token")
	// 需要登陆
	if _token != "" {
		//var user *user2.Model
		//result := db.DefaultDB.Table("users").Where("token = ?", _token).First(&user)
		//if result != nil {
		//	fmt.Println(user)
		//	fmt.Println("验证通过")
		//}
		next.Next()
	}else {
		fmt.Println( "isHave(next.Request.URL.Path)" ,isHave(next.Request.URL.Path))
		if isHave(next.Request.URL.Path) {
			next.Next()
			return
		}
		next.JSON(http.StatusBadGateway, models.HttpModels{
			Code: config.HttpNeedLogin,
			Message: config.HttpNeedLoginMessage,
		})
		next.Abort()
	}
}
