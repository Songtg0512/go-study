package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/banner?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	fmt.Println("数据库连接成功！")

	db.AutoMigrate(&Student{})

	students := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}

	// 添加用户
	db.Create(&students)

	// 查询年龄 > 18 岁的学生
	var student Student
	db.Model(&students).Where("age > ?", 18).First(&student)
	fmt.Println(student)

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	// 查询改正后的
	var student2 Student
	db.Model(&Student{}).Where("name = ?", "张三").First(&student2)
	fmt.Println(student2)

	db.Where("age < ?", 15).Delete(&Student{})
}
