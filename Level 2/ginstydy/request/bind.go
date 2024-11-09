package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func mainb() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		// // 查询参数
		// type User struct {
		// 	Name string `form:"name"`
		// 	Age  int    `form:"age"`
		// }
		// var user User
		// err := c.ShouldBindQuery(&user) //类型不对会报错
		// fmt.Println(user, err)

		// // 路径参数
		// r.GET("users/:id/:name", func(c *gin.Context) {
		// 	type User struct {
		// 		Name string `uri:"name"`
		// 		Id   int    `uri:"id"`
		// 	}
		// 	var user User
		// 	err := c.ShouldBindUri(&user)
		// 	fmt.Println(user, err)
		// })

		// // 表单参数 不能解析x-www-form-urlencoded的格式
		// r.POST("form", func(c *gin.Context) {
		// 	type User struct {
		// 		Name string `form:"name"`
		// 		Age  int    `form:"age"`
		// 	}
		// 	var user User
		// 	err := c.ShouldBind(&user)
		// 	fmt.Println(user, err)
		// })

		// // json参数
		// r.POST("json", func(c *gin.Context) {
		// 	type User struct {
		// 		Name string `json:"name"`
		// 		Age  int    `json:"age"`
		// 	}
		// 	var user User
		// 	err := c.ShouldBindJSON(&user)
		// 	fmt.Println(user, err)
		// })

		// header参数
		r.POST("/header", func(c *gin.Context) {
			type User struct {
				Name        string `header:"Name"`
				Age         int    `header:"Age"`
				UserAgent   string `header:"User-Agent"`
				ContentType string `header:"Content-Type"`
			}
			var user User
			err := c.ShouldBindHeader(&user)
			fmt.Println(user, err)
		})

	})

	r.Run(":8080")
}
