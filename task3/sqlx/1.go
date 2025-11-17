package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}

	// ---------- 题目1 ----------
	var techEmployees []Employee
	err = db.Select(&techEmployees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("技术部员工：", techEmployees)

	var topEmployee Employee
	err = db.Get(&topEmployee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("工资最高的员工：", topEmployee)

	// ---------- 题目2 ----------
	var expensiveBooks []Book
	err = db.Select(&expensiveBooks, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("价格大于50元的书籍：", expensiveBooks)
}
