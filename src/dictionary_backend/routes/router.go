package routes

import (
	"dictionary_backend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/get/:language/:str", controller.GetInfoByName)
	router.POST("/upload/:language/:str/:sign", controller.PostInfoByName)
	return router
}
