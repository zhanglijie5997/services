package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	"service/models"
)

type Data struct {
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
}

func User(context *gin.Context)  {
	context.JSON(http.StatusOK, models.HttpModels{
		Code: config.HttpSuccess,
		Message: config.HttpSuccessMessage,
		Data: Data{
			Name: "zlj",
			Age: 12,
			Sex: 12,
		},
	})
}