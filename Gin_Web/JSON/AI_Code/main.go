package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体，用来“接”JSON数据
type User struct {
	// 后面这个 `json:"name"` 叫做 Tag（标签）
	// 意思是：JSON里的 "name" 字段，对应我也在这个 Name 字段里
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 1. 模拟收到一段 JSON 文本（比如前端发来的）
	jsonStr := `{"name": "Ken", "age": 25}`

	var user User
	// 2. 反序列化：把 JSON 变成 Go 结构体
	// &user 表示把数据直接写进 user 变量的内存地址里
	json.Unmarshal([]byte(jsonStr), &user)

	fmt.Println(user.Name) // 输出: Ken

	// 3. 序列化：把 Go 结构体变成 JSON
	user.Age = 26
	newJson, _ := json.Marshal(user)
	fmt.Println(string(newJson)) // 输出: {"name":"Ken","age":26}
}
