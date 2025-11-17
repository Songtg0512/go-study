package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 账户表
type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance float64
}

// 交易表
type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	fmt.Println("数据库连接成功！")

	// 自动建表
	if err := db.AutoMigrate(&Account{}, &Transaction{}); err != nil {
		panic("自动建表失败：" + err.Error())
	}

	// 初始化两个账户
	db.Create(&Account{ID: 1, Balance: 500})
	db.Create(&Account{ID: 2, Balance: 200})

	// 转账金额
	amount := 100.0

	// 使用事务执行转账
	err = db.Transaction(func(tx *gorm.DB) error {
		var from Account
		var to Account

		// 检查余额
		if from.Balance < amount {
			return fmt.Errorf("账户余额不足")
		}

		// 扣除账户 A
		from.Balance -= amount
		if err := tx.Save(&from).Error; err != nil {
			return err
		}

		to.Balance += amount
		if err := tx.Save(&to).Error; err != nil {
			return err
		}

		// 记录交易
		transaction := Transaction{
			FromAccountID: from.ID,
			ToAccountID:   to.ID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	if err != nil {
		fmt.Println("转账失败，已回滚：", err)
	} else {
		fmt.Println("转账成功！")
	}

	// 查询账户余额
	var accounts []Account
	db.Find(&accounts)
	fmt.Println("账户余额：", accounts)
}
