package main

import "github.com/gin-gonic/gin"

func mainf() {
	r := gin.Default()

	r.GET("", func(ctx *gin.Context) {
		// 设置Content唤起浏览器下载  只能是get请求
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; filename=file.go")
		ctx.File("file.go")

		//前端请求后端接口下载
		//返回 blob

		//最好的做法是前端请求接口，后端返回一个临时下载地址
	})

	r.Run(":8080")
}
