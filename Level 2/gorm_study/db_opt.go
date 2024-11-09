package main

import (
	"fmt"
	"studyGorm/global"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `gorm:"primaryKey"`      // 主键
	Name      string `gorm:"not null;unique"` // 不能为空，且唯一
	Age       int    `gorm:"default:18"`
	Email     string
	CreatedAt time.Time // 在创建记录时自动设置为当前时间
	UpdatedAt time.Time // 在创建记录时自动设置为当前时间
}

func Age18(tx *gorm.DB) *gorm.DB {
	return tx.Where("age >= ?", 18)
}

func insert() {

	// 插入数据
	// err := global.DB.Create(&User{
	// 	Name: "龙龙",
	// 	Age:  25,
	// }).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // 回填式插入数据
	// user := User{
	// 	Name: "龙龙1",
	// 	Age:  16,
	// }
	// err := global.DB.Create(&user).Error
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(user.ID, user.Name, user.Age)

	// 批量 插入数据
	var users = []User{
		{
			Name: "王五",
			Age:  16,
		},
		{
			Name: "李四",
			Age:  16,
		},
	}
	err := global.DB.Create(&users).Error
	fmt.Println(users, err)
}

func query() {
	// 查全部
	var users []User
	global.DB.Find(&users)
	fmt.Println(users)

	// 带条件查询
	users = []User{}
	global.DB.Find(&users, "name = ?", "张三")
	fmt.Println(users)

	// 获取一条记录
	var user User
	global.DB.Take(&user)
	fmt.Println(user)

	// 如果查不到则会报错
	user = User{}
	err := global.DB.Take(&user, "id = ?", 100).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("不存在的记录")
	}

	// 如果不想出现错误
	user = User{}
	err = global.DB.Limit(1).Find(&user, "id = ?", 100).Error
	fmt.Println(err)

	// 打印实际的sql
	global.DB.Debug().Take(&user, "id = ?", 1)
}

// 更新数据
func update() {
	// 有很多方法，Save、Update、UpdateColumn、Updates

	// 不同的方法有不同的区别，如下：

	// Save，有主键记录就是更新，并且可以更新零值，否则就是创建
	// Update，可以更新零值，必须要有条件
	// UpdateColumn，可以更新零值，不会走更新的Hook
	// Updates，如果是结构体，则更新非零值，map可以更新零值

	// // 创建
	// var user = User{
	// 	Name: "枫枫1",
	// }
	// global.DB.Save(&user)
	// fmt.Println(user)

	// // 更新
	// user = User{
	// 	ID:        9,
	// 	Name:      "枫枫2",
	// 	CreatedAt: time.Now(),
	// }
	// global.DB.Save(&user)
	// fmt.Println(user)

	// // Update和UpdateColumn
	// global.DB.Model(&User{}).Where("id = ?", 1).Update("name", "张三") // 走了更新钩子

	// global.DB.Model(&User{}).Where("id = ?", 1).UpdateColumn("name", "张三")

	// // Updates
	// var user = User{ID: 1}

	// global.DB.Model(&user).Updates(User{
	// 	Name: "张三丰",
	// })

	// user = User{ID: 2}
	// // 不会更新零值
	// global.DB.Model(&user).Updates(User{
	// 	Name: "",
	// })

	// user = User{ID: 2}
	// // 会更新零值
	// global.DB.Model(&user).Updates(map[string]any{
	// 	"name": "",
	// })

	// // Expr
	// // 通常用于获取原字段的数据
	// // 例如年龄加一
	// global.DB.Model(&User{}).Where("id = ?", 1).UpdateColumn("age", gorm.Expr("age + 1"))

	// 删除
	var user = User{ID: 10}
	global.DB.Delete(&user)
	global.DB.Delete(&User{}, 9)
	// 批量删除
	global.DB.Delete(&User{}, []int{1, 2, 3})
}

// 高级查询
func HightQuery() {
	// var user User
	// global.DB.Where("age > ?", 20).Take(&user)
	// fmt.Println(user)
	// // 注意where的顺序，再&函数的前面

	// var user = User{ID: 5, Name: "李四"}
	// global.DB.Debug().Where(user).Take(&user)
	// // SELECT * FROM `user_models` WHERE (`user_models`.`id` = 5 AND `user_models`.`name` = '李四') AND `user_models`.`id` = 5IMIT 1
	// fmt.Println(user)

	// var user User
	// global.DB.Debug().Where(map[string]any{"age": 0,}).Take(&user)
	// // SELECT * FROM `user_models` WHERE `age` = 0  LIMIT 1
	// fmt.Println(user)

	// query := global.DB.Where("age = ? and name = ?", 11, "王五1")
	// var user User
	// global.DB.Debug().Where(query).Take(&user)
	// // SELECT * FROM `user_models` WHERE (age = 11 and name = '王五1') LIMIT 1
	// fmt.Println(user)

	// var user User
	// global.DB.Debug().Or("name = ?", "王五").Or("age = 12").Take(&user)
	// // SELECT * FROM `user_models` WHERE (name = '王五' OR age = 12) LIMIT 1
	// fmt.Println(user)

	// var user User
	// global.DB.Debug().Not("age = 12").Take(&user)
	// //  SELECT * FROM `user_models` WHERE NOT age = 12 LIMIT 1
	// fmt.Println(user)

	// var userList []User
	// // 降序
	// global.DB.Order("age desc").Find(&userList)
	// fmt.Println(userList)
	// // 升序
	// global.DB.Order("age asc").Find(&userList)
	// fmt.Println(userList)

	// Scan
	// var nameList []string
	// global.DB.Model(User{}).Select("name").Scan(&nameList)
	// fmt.Println(nameList)

	// var nameList1 []string
	// global.DB.Model(User{}).Pluck("name", &nameList)
	// fmt.Println(nameList1)

	// // 分页
	// var users []User
	// // // 第一页
	// // global.DB.Limit(10).Offset(0).Find(&users)
	// // fmt.Println(users)
	// // // 第二页
	// // global.DB.Limit(10).Offset(10).Find(&users)
	// // fmt.Println(users)
	// // // 第三页
	// // global.DB.Limit(10).Offset(20).Find(&users)
	// // fmt.Println(users)
	// // 第n页
	// limit := 2
	// page := 2
	// global.DB.Limit(limit).Offset((page - 1) * limit).Find(&users)
	// fmt.Println(users)
}

func main() {
	global.Connect()
	// insert()
	// query()
	// update()
	HightQuery()

	// var users []User
	// global.DB.Scopes(Age18).Find(&users)
	// fmt.Println(users)
}
