package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ---------------------------
// 题目1：模型定义
// ---------------------------

// User 用户模型
type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	ArticleCount int    // 文章数量统计
	Posts        []Post `gorm:"foreignKey:UserID"`
}

// Post 文章模型
type Post struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Content      string
	UserID       uint
	Comments     []Comment `gorm:"foreignKey:PostID"`
	CommentCount int       // 评论数量统计
}

// Comment 评论模型
type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	PostID  uint
}

// ---------------------------
// 题目3：钩子函数
// ---------------------------

// GORM 钩子函数：Post 在创建后更新用户文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 增加用户文章数量
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
}

// Comment 删除后检查文章评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", 0).Error
	} else {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", count).Error
	}
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	// 自动建表
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		panic("自动建表失败：" + err.Error())
	}

	fmt.Println("数据库表创建完成！")

	// ---------------------------
	// 插入测试数据
	// ---------------------------
	user := User{Name: "张三"}
	db.Create(&user)

	post1 := Post{Title: "文章一", Content: "内容一", UserID: user.ID}
	post2 := Post{Title: "文章二", Content: "内容二", UserID: user.ID}
	db.Create(&post1)
	db.Create(&post2)

	comment1 := Comment{Content: "评论1", PostID: post1.ID}
	comment2 := Comment{Content: "评论2", PostID: post1.ID}
	comment3 := Comment{Content: "评论3", PostID: post2.ID}
	db.Create(&comment1)
	db.Create(&comment2)
	db.Create(&comment3)

	// 更新每篇文章的评论数量统计
	db.Model(&Post{}).Where("id = ?", post1.ID).Update("comment_count", 2)
	db.Model(&Post{}).Where("id = ?", post2.ID).Update("comment_count", 1)

	// ---------------------------
	// 题目2：关联查询
	// ---------------------------

	// 查询某个用户的所有文章及其评论
	var postsWithComments []Post
	db.Preload("Comments").Where("user_id = ?", user.ID).Find(&postsWithComments)
	fmt.Println("用户文章及评论：")
	for _, p := range postsWithComments {
		fmt.Printf("文章: %s, 评论数量: %d\n", p.Title, len(p.Comments))
		for _, c := range p.Comments {
			fmt.Printf("   评论: %s\n", c.Content)
		}
	}

	// 查询评论数量最多的文章
	var mostCommentedPost Post
	db.Order("comment_count DESC").First(&mostCommentedPost)
	fmt.Printf("评论最多的文章: %s, 评论数量: %d\n", mostCommentedPost.Title, mostCommentedPost.CommentCount)

	// ---------------------------
	// 测试删除评论触发钩子
	// ---------------------------
	db.Delete(&comment1) // 删除评论1
	var updatedPost Post
	db.First(&updatedPost, post1.ID)
	fmt.Printf("删除评论后文章评论数量: %d\n", updatedPost.CommentCount)
}
