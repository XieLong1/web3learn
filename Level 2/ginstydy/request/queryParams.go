package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func mainq() {
	r := gin.Default()

	// 查询参数
	// r.GET("", func(c *gin.Context) {
	// 	name := c.Query("name")
	// 	age := c.DefaultQuery("age", "25")
	// 	keyList := c.QueryArray("key")
	// 	fmt.Println(name, age, keyList)
	// })
	// http://127.0.0.1:8080/?name=longlong&age=30&key=123&key=456
	// 结果	longlong 30 [123 456]

	// 动态参数
	// r.GET("users/:id/:name", func(c *gin.Context) {
	// 	userId := c.Param("id")
	// 	name := c.Param("name")
	// 	fmt.Println(userId, name)
	// })
	// http://127.0.0.1:8080/users/123/longlong
	// 输出 123 longlong

	// 表单参数，一般就是from表单 /from-data
	r.POST("user", func(c *gin.Context) {
		// name := c.PostForm("name")
		// age, ok := c.GetPostForm("age")
		// fmt.Println(name, age, ok)
		// ok 可以判断参数前端传了还是没传

		// // 文件上传 单文件
		// fileHeader, err := c.FormFile("file")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println(fileHeader.Filename) //文件名称
		// fmt.Println(fileHeader.Size)     // 文件大小 单位是字节

		// // // 传统方式保存
		// // file, _ := fileHeader.Open()
		// // byteData, _ := io.ReadAll(file)
		// // err = os.WriteFile("xxx.jpg", byteData, 0666)
		// // fmt.Println(err)

		// // // 快捷方式
		// // err = c.SaveUploadedFile(fileHeader, "update/xxx/"+fileHeader.Filename) // 如果路径不存在就会创建
		// // fmt.Println(err)

		// 多文件上传

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 第一个参数是file可以忽略 第二个参数是header数组
		for _, headers := range form.File {
			for _, header := range headers {
				c.SaveUploadedFile(header, "updates/xxx"+header.Filename)
			}
		}

	})

	r.Run(":8080")
}
