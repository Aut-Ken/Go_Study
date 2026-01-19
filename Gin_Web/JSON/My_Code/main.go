package main

import (
	"encoding/json"
	"fmt"
)

// 1. 定义结构体
type Book struct {
	// TODO: 在这里定义三个字段：
	// Title (string)
	// Price (float64)
	// OnSale (bool)
	// 重要：一定要加上 `json:"..."` 标签，让输出的 JSON key 变成全小写！
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	OnSale bool    `json:"onsale"`
}

func main() {
	// --- 目标 A: 把 Go 变成 JSON ---
	// 2. 创建一个书的对象
	myBook := Book{
		// TODO: 初始化这本书：标题是 "Go Action", 价格 59.9, 在售 true
		Title:  "Go Action",
		Price:  59.9,
		OnSale: true,
	}

	// 3. 将 myBook 转换成 JSON
	// 提示：使用 json.Marshal(myBook)
	// jsonBytes, err := ...
	jsonBytes, _ := json.Marshal(myBook)
	fmt.Println(string(jsonBytes))
	fmt.Println(jsonBytes)
	// TODO: 打印转换后的 JSON 字符串
	// fmt.Println(string(jsonBytes))

	// --- 目标 B: 把 JSON 变成 Go ---
	// 这是一个前端传来的 JSON 字符串
	incomingJson := `{"title": "Harry Potter", "price": 99.9, "onsale": false}`

	var anotherBook Book

	// 4. 将 incomingJson 解析到 anotherBook 变量里
	// 提示：使用 json.Unmarshal(...)
	// 注意：Unmarshal 的第一个参数必须是 []byte 类型，第二个参数必须是 &指针
	json.Unmarshal([]byte(incomingJson), &anotherBook)
	fmt.Printf("收到新书: %s\n 价格是: %.2f\n", anotherBook.Title, anotherBook.Price)
	// TODO: 打印解析出来的书名 (Title)
	// fmt.Printf("收到新书：%s\n", ...)
}
