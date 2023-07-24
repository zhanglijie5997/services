package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	"service/models"
)

func NoRoute(context *gin.Context)  {
	context.JSON(http.StatusOK, models.HttpModels{
		Code: config.HttpNotRoute,
		Message: config.HttpNotRouteMessage,
	})
}