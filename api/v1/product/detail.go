package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/config"
	"service/models"
)

func Detail(context *gin.Context) {
	_id := context.DefaultQuery("id", "")
	if _id == "" {
		context.JSON(http.StatusOK, config.HttpCodeMesaage[config.HttpNeedId])
		return
	}
	//_detail := core.DB.Where("")
	fmt.Println(_id, "_id111")
	context.JSON(http.StatusOK, models.HttpModels{
		Code:    config.HttpSuccess,
		Message: config.HttpSuccessMessage,
	})
}
