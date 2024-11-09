package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// st是别名，st的路由都会到static文件夹下面
	r.Static("st", "static")
	// 访问具体的稳健
	r.StaticFile("abc", "static/abc.txt")
	// 静态文件路径不能再被路由使用了

	r.Run(":8080")
}
