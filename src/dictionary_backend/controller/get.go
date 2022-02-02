package controller

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"dictionary_backend/configure"

	"github.com/gin-gonic/gin"
)

type Location struct {
	sign    string `db:"sign"`
	address string `db:"address"`
}

type Pronounce struct {
	sign string `db:"sign"`
	str  string `db:"str"`
}

var Db *sqlx.DB

func GetInfoByName(c *gin.Context) {
	str := c.Param("str")
	fmt.Println(str)
	ConnectSentence := configure.MysqlUsername + ":" + configure.MysqlPassword + "@tcp(" + configure.Address + ":" + configure.MysqlPort + ")/" + configure.DatabaseName
	database, err1 := sqlx.Connect("mysql", ConnectSentence)
	if err1 != nil {
		fmt.Println("open mysql failed,", err1)
		return
	}

	Db = database

	rows, err2 := Db.Query("select sign from pronounce where str = '" + str + "'")
	if err2 != nil {
		fmt.Println("exec failed, ", err2)
		return
	}

	for rows.Next() {
		var sign string
		var strr string

		rows.Scan(&sign, &strr)

		fmt.Println(sign)
		fmt.Println(strr)
	}

	fmt.Println("select success:", rows)
	//	c.String(http.StatusOK, pronounce)
	Db.Close() // 注意这行代码要写在上面err判断的下面
}
