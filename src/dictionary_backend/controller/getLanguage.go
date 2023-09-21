package controller

import (
	"dictionary_backend/configure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func GetGroup(c *gin.Context) {
	column := c.Param("column")

	database, err1 := sqlx.Connect("mysql", configure.ConnectSentence)
	if err1 != nil {
		fmt.Println("open mysql failed,", err1)
		return
	}

	configure.Db = database

	rows, err2 := configure.Db.Query("select " + column + " from pronounce group by " + column)
	if err2 != nil {
		fmt.Println("exec failed, ", err2)
		return
	}

	var lan string
	var Array []string
	for rows.Next() {
		err3 := rows.Scan(&lan)
		if err3 != nil {
			fmt.Println("scan failed, ", err3)
			return
		}
		Array = append(Array, lan)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Array,
	})

	err4 := configure.Db.Close()
	if err4 != nil {
		fmt.Println("close failed, ", err4)
		return
	}
}
