package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	DB, err := sql.Open("mysql", "root:123123@tcp(localhost:3306)/test")

	// 设置数据库的最大连接数
	DB.SetConnMaxLifetime(100)

	// 关闭连接
	defer DB.Close()

	// 最大闲置数
	DB.SetMaxIdleConns(10)

	fmt.Errorf("%d", err)

	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("连接成功")

	result, _ := DB.Exec("SELECT * FROM tb_admin")

	fmt.Println(result)

}
