package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	"service/models"
	user2 "service/models/user"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

func Home(context *gin.Context)  {
	user := &user2.User{Sex: 1, Age: 12, UserName: "zlj"}
	data := models.DATA{
		Data: user,
	}
	context.JSON(http.StatusOK, models.HttpModels{
		Code: config.HttpSuccess,
		Message: config.HttpSuccessMessage,
		Data: data.Data,
	})

}