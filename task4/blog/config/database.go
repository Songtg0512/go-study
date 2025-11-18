package config

import (
	"log"

	"blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	// 使用 SQLite 数据库，方便开发和测试
	DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移数据库表结构
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
}
