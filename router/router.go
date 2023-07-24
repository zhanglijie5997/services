package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"service/handler"
	"service/middleware"
	v1 "service/router/v1"
)
var Runner *gin.Engine

func Router()  {
	Runner = gin.Default()
	Runner.NoMethod(handler.NoMethod)
	Runner.NoRoute(handler.NoRoute)
	Runner.Use(cors.Default())
	Runner.Use(middleware.Middlewares)
	v1.RouterGroupV1(Runner)
	err := Runner.Run(":9000")
	if err != nil {
		return 
	}
}