package routes

import (
	"dictionary_backend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getSign/:language/:str", controller.GetSignAndLocationByStr)
	router.GET("/getAudio", controller.GetAudioByLocation)
	router.POST("/upload/:language/:str/:sign", controller.PostInfo)
	router.GET("/getGroup/:column", controller.GetGroup)
	return router
}
