package controller

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"dictionary_backend/configure"

	"github.com/gin-gonic/gin"
)

var Db *sqlx.DB

func GetInfoByName(c *gin.Context) {
	language := c.Param("language")
	str := c.Param("str")
	ConnectSentence := configure.MysqlUsername + ":" + configure.MysqlPassword + "@tcp(" + configure.Address + ":" + configure.MysqlPort + ")/" + configure.DatabaseName
	database, err1 := sqlx.Connect("mysql", ConnectSentence)
	if err1 != nil {
		fmt.Println("open mysql failed,", err1)
		return
	}

	Db = database

	rows, err2 := Db.Query("select * from pronounce where str = '" + str + "' and language = '" + language + "'")
	if err2 != nil {
		fmt.Println("exec failed, ", err2)
		return
	}

	var data []string
	for rows.Next() {
		var sign string
		var language string
		var word string

		err3 := rows.Scan(&sign, &language, &word)
		if err3 != nil {
			fmt.Println("scan failed, ", err3)
			return
		}

		data = append(data, sign)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})

	//	c.String(http.StatusOK, pronounce)
	err4 := Db.Close()
	if err4 != nil {
		fmt.Println("close failed, ", err4)
		return
	} // 注意这行代码要写在上面err判断的下面
}
