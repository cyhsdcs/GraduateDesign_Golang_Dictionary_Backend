package configure

import "github.com/jmoiron/sqlx"

var Address = "localhost" //服务器地址

var MysqlUsername = "root"        //MySQL用户名
var MysqlPassword = "Apolo7g8ies" //MySQL密码
var MysqlPort = "3306"            //MySQL端口
var DatabaseName = "dictionary"   //数据库名
var TableName = "pronounce"       //表名
var ConnectSentence = MysqlUsername + ":" + MysqlPassword + "@tcp(" + Address + ":" + MysqlPort + ")/" + DatabaseName
var Db *sqlx.DB

//var Path = "../../../music"

var Path = "./music"
