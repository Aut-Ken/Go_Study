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
	// ==========================================
	/*
		fmt.Println("--- 开始进行寿命转移仪式 ---")

		// 1. 开启事务 (注意：这里返回的是 tx 对象，不是 db 对象了！)
		tx, err := db.Begin()
		if err != nil {
			log.Fatal("开启事务失败: ", err)
		}

		// 准备一个“后悔药”机制
		// defer 的作用是：函数退出前，不管成功失败，都检查一下。
		// 如果没 Commit，就自动 Rollback (防止程序崩了导致死锁)
		defer tx.Rollback()

		// 2. 第一步：赵艺凯 (id=1) 减 10 岁
		// 注意：这里用的是 tx.Exec，不是 db.Exec ！！
		_, err = tx.Exec("UPDATE users SET age = age - 10 WHERE id = ?", 1)
		if err != nil {
			// 如果这一步错了，defer 会自动触发 Rollback
			log.Fatal("扣减寿命失败: ", err)
		}

		// =============================================
		// 🔥 模拟一个意外！
		// 假设这里突然停电了，或者代码写错了 (比如故意写错表名 usersss)
		// =============================================
		_, err = tx.Exec("UPDATE users SET age = age + 10 WHERE id = ?", 2)
		if err != nil {
			// 报错了！此时程序会退出，defer 里的 Rollback 会执行
			// 数据库会发现：刚才赵艺凯减掉的 10 岁会瞬间恢复！
			log.Println("❌ 仪式失败！对方接收失败，开始回滚...")
			return
		}

		// 3. 如果上面都没报错，提交事务！
		err = tx.Commit()
		if err != nil {
			log.Fatal("提交失败: ", err)
		}

		fmt.Println("✅ 仪式完成！交易生效。")
	*/
}
