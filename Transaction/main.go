package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// 1. 准备连接信息 (DSN)
	// 格式说明: 用户名:密码@tcp(IP地址:端口)/数据库名
	dsn := "root:@tcp(127.0.0.1:3306)/go_db"

	// 2. 初始化连接对象
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接配置写错了: ", err)
	}
	defer db.Close()

	// 3. 测试连接 (Ping)
	err = db.Ping()
	if err != nil {
		log.Fatal("敲门失败，可能是 MySQL 没启动或密码错了: ", err)
	}
	fmt.Println("🎉 恭喜！数据库连接成功！")
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("查询指令发送失败: ", err)
	}
	defer rows.Close() // 用完也要记得关闭结果集

	fmt.Println("-------------------------")

	// 4. 结果测试
	columns, _ := rows.Columns()
	count := len(columns)

	values := make([]interface{}, count)
	valuesPtrs := make([]interface{}, count)

	for rows.Next() {
		for i := range columns {
			valuesPtrs[i] = &values[i]
		}

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			log.Fatal(err)
		}

		for i, colName := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				val = string(b)
			}
			fmt.Printf("%s: %v | ", colName, val)
		}
		fmt.Println()
	}

	fmt.Println("-------------------------")

	// 5. 添加数据
	/*

		newUserName := "路人甲"
		newUserAge := 30

		result, err := db.Exec("INSERT INTO users (id, name, age) VALUES (?, ?, ?)", 2, newUserName, newUserAge)
		if err != nil {
			log.Fatal("插入数据失败：", err)
		}

		lastID, _ := result.LastInsertId()
		fmt.Printf("✨ 成功添加新用户！分配的 ID 是: %d\n", lastID)
	*/

	// 6. Updata
	/*
		result, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", "陆奕涵(升级版)", 22, 2)
		if err != nil {
			log.Fatal("更新失败: ", err)
		}

		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("✅ 成功修改了 %d 行数据\n", rowsAffected)
	*/

	// 7. Delete
	/*
		delResult, err := db.Exec("DELETE FROM users WHERE id = ?", 2)
		if err != nil {
			log.Fatal("更新失败: ", err)
		}
		delRows, _ := delResult.RowsAffected()
		fmt.Printf("🗑️ 成功删除了 %d 行数据\n", delRows)
	*/

	// 8. Transcation
	fmt.Println("--- 开始进行寿命转移仪式 ---")

}
