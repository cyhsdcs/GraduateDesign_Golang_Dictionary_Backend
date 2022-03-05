package controller

import (
	"dictionary_backend/configure"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
)
import "fmt"

func PostInfo(c *gin.Context) {
	language := c.Param("language")
	str := c.Param("str")
	sign := c.Param("sign")

	file, _ := c.FormFile("file")

	path := configure.Path + "/" + language + "/" + str + "/" + sign + "/"
	//如果没有path文件目录就创建一个
	if _, err := os.Stat(path); err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Println("mkdir failed! ", err)
				return
			}
		}
	}

	err1 := c.SaveUploadedFile(file, path+file.Filename)
	if err1 != nil {
		fmt.Println("save file failed! ", err1)
		return
	}

	database, err2 := sqlx.Connect("mysql", configure.ConnectSentence)
	if err2 != nil {
		fmt.Println("open mysql failed,", err2)
		return
	}

	configure.Db = database

	r, err3 := configure.Db.Exec("insert into "+configure.TableName+" (language, str, sign, location)values(?, ?, ?, ?)", language, str, sign, configure.Path+file.Filename)
	if err3 != nil {
		fmt.Println("exec failed, ", err3)
		return
	}
	id, err4 := r.LastInsertId()
	if err4 != nil {
		fmt.Println("exec failed, ", err4)
		return
	}

	fmt.Println("insert success! ", id)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
