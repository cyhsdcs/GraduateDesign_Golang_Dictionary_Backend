package routes

import (
	"dictionary_backend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/get/:str", controller.GetInfoByName)

	return router
}
