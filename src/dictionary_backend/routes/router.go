package routes

import (
	"dictionary_backend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/get/:language/:str", controller.GetSignAndLocationByStr)
	router.POST("/upload/:language/:str/:sign", controller.PostInfo)
	return router
}
