package main

import (
	"github.com/gin-gonic/gin"
	"study.com/project/response/res"
)

func main1() {
	gin.SetMode("release")
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		res.OkWithMsg(c, "登录成功")
	})

	r.GET("/users", func(c *gin.Context) {
		res.OkWithData(c, gin.H{"name": "龙龙"})
	})

	r.POST("/users", func(c *gin.Context) {
		res.FailWithMsg(c, "参数错误")
	})
	r.Run(":8080")
}
