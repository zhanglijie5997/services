package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	"service/models"
)

func NoMethod(context *gin.Context)  {
	context.JSON(http.StatusOK, models.HttpModels{
		Code: config.HttpNotMethod,
		Message: config.HttpNotMethodMessage,
	})
}
