package main

import "github.com/gin-gonic/gin"

func main2() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", map[string]any{
			"title": "leon",
		})
	})

	r.Run(":8080")
}
