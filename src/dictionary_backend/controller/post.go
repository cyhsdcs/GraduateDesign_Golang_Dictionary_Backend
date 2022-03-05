package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "fmt"

func PostInfoByName(c *gin.Context) {
	file, _ := c.FormFile("file")

	err := c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
