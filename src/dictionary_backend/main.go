package main

//导入包
import (
	"dictionary_backend/routes"
)

func main() {
	//初始化引擎实例
	router := routes.InitRouter()
	//注册一个Get请求的方法

	router.Run(":8888") //默认8080端口 自己指定例如： router.Run(":8888")
}
