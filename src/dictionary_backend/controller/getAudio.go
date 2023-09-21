package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

func GetAudioByLocation(c *gin.Context) {
	location := c.Query("location")
	temp, errByOpenFile := os.Open(location)
	defer temp.Close()
	fileName := path.Base(location)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; path="+fileName)

	//非空处理
	if errByOpenFile != nil {
		c.Redirect(http.StatusFound, errByOpenFile.Error())
		return
	}
	//c.String(http.StatusOK, location)

	c.File(location)
	return
}
