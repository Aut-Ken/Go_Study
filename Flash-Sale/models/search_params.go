package models

// UserSearchParams 专门用来接收前端的查询条件
type UserSearchParams struct {
	Name   *string `json:"name"`    // 用指针，nil 代表没传，"" 代表搜空字符串
	MinAge *int    `json:"min_age"` // 最小年龄
	MaxAge *int    `json:"max_age"` // 最大年龄
	ID     *uint   `json:"id"`      // 精确 ID
}
