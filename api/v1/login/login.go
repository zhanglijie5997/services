package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	"service/db/user"
	"service/models"
)


func Login(context *gin.Context)  {
	token := context.GetHeader("token")
	fmt.Println("user token", token)
	context.JSON(http.StatusOK, models.HttpModels{
		Code: config.HttpSuccess,
		Message: config.HttpSuccessMessage,
		Data: user.User{},
	})
}
