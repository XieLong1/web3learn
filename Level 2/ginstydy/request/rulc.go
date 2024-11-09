package main

import (
	"github.com/gin-gonic/gin"
)

// 常用验证器

// 不能为空，并且不能没有这个字段
// required： 必填字段，如：binding:"required"

// // 针对字符串的长度
// min 最小长度，如：binding:"min=5"
// max 最大长度，如：binding:"max=10"
// len 长度，如：binding:"len=6"

// // 针对数字的大小
// eq 等于，如：binding:"eq=3"
// ne 不等于，如：binding:"ne=12"
// gt 大于，如：binding:"gt=10"
// gte 大于等于，如：binding:"gte=10"
// lt 小于，如：binding:"lt=10"
// lte 小于等于，如：binding:"lte=10"

// // 针对同级字段的
// eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
// nefield 不等于其他字段的值

// - 忽略字段，如：binding:"-"

// gin内置验证器

// 枚举  只能是red 或green
// oneof=red green

// // 字符串
// contains=fengfeng  // 包含fengfeng的字符串
// excludes // 不包含
// startswith  // 字符串前缀
// endswith  // 字符串后缀

// // 数组
// dive  // dive后面的验证就是针对数组中的每一个元素

// // 网络验证
// ip
// ipv4
// ipv6
// uri
// url
// uri 在于I(Identifier)是统一资源标示符，可以唯一标识一个资源。
// url 在于Locater，是统一资源定位符，提供找到该资源的确切路径

// 日期验证  1月2号下午3点4分5秒在2006年
// datetime=2006-01-02

// 自定义验证的错误信息

// type UserInfo struct {
//   Username string `json:"username" binding:"required" msg:"用户名不能为空"`
//   Password string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
//   Email    string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
// }

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		// 参数校验
		// json参数
		r.POST("json", func(c *gin.Context) {
			type User struct {
				Name string `json:"name" binding:"required,min=3,max=5"`
				Age  int    `json:"age"`
			}
			var user User
			err := c.ShouldBindJSON(&user)
			if err != nil {
				c.String(200, err.Error())
				return
			}
			c.JSON(200, user)
		})

	})

	r.Run(":8080")
}
