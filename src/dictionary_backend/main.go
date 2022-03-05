package main

//导入包
import (
	"dictionary_backend/routes"
)

func main() {
	//初始化引擎实例
	router := routes.InitRouter()

	err := router.Run(":8888")
	if err != nil {
		return
	} //默认8080端口 自己指定例如： router.Run(":8888")
}
