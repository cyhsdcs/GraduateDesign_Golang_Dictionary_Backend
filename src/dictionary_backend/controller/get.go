package controller

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"dictionary_backend/configure"

	"github.com/gin-gonic/gin"
)

func GetSignAndLocationByStr(c *gin.Context) {
	language := c.Param("language")
	str := c.Param("str")

	database, err1 := sqlx.Connect("mysql", configure.ConnectSentence)
	if err1 != nil {
		fmt.Println("open mysql failed,", err1)
		return
	}

	configure.Db = database

	rows, err2 := configure.Db.Query("select * from pronounce where str = '" + str + "' and language = '" + language + "'")
	if err2 != nil {
		fmt.Println("exec failed, ", err2)
		return
	}

	var data map[string][]string
	data = make(map[string][]string)
	var lan string
	var word string
	var sign string
	var location string
	lastSign := ""

	var locationArray []string
	for rows.Next() {

		err3 := rows.Scan(&lan, &word, &sign, &location)
		if err3 != nil {
			fmt.Println("scan failed, ", err3)
			return
		}
		if sign != lastSign {
			if locationArray != nil {
				data[lastSign] = locationArray
				locationArray = nil
			}
			lastSign = sign
		}
		locationArray = append(locationArray, location)
	}
	data[lastSign] = locationArray

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})

	//	c.String(http.StatusOK, pronounce)
	err4 := configure.Db.Close()
	if err4 != nil {
		fmt.Println("close failed, ", err4)
		return
	} // 注意这行代码要写在上面err判断的下面
}
