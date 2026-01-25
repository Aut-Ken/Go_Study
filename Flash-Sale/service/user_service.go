package service

import (
	"errors"
	"flash-sale/database"
	models "flash-sale/models"

	"gorm.io/gorm"
)

// TransferLife 寿命转移的核心业务逻辑
// 入参：fromID, toID
// 返回：error (如果成功返回 nil)
func TransferLife(fromID uint, toID uint) error {
	// 开启事务
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 扣减 (使用 gorm.Expr 原子更新)
		// 注意：这里用 tx，不要用 database.DB
		result := tx.Model(&models.User{}).Where("id = ?", fromID).
			Update("age", gorm.Expr("age - ?", 10))

		if result.Error != nil {
			return result.Error
		}
		// 检查有没有扣减成功 (防止 ID 不存在)
		if result.RowsAffected == 0 {
			return errors.New("扣减方 ID 不存在")
		}

		// 2. 增加
		result = tx.Model(&models.User{}).Where("id = ?", toID).
			Update("age", gorm.Expr("age + ?", 10))

		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("接收方 ID 不存在")
		}

		return nil // 提交事务
	})
}

// GetUsers 动态查询用户
func GetUsers(params models.UserSearchParams) ([]models.User, error) {
	var users []models.User

	// 1. 拿到一个基础的数据库连接（还没开始查）
	// 相当于写了: SELECT * FROM users
	query := database.DB.Model(&models.User{})

	// 2. 像搭积木一样，根据条件动态添加 Where 子句

	// 如果传了 Name，就模糊查询
	if params.Name != nil {
		query = query.Where("name LIKE ?", "%"+*params.Name+"%")
	}

	// 如果传了 ID，就精确查询
	if params.ID != nil {
		query = query.Where("id = ?", *params.ID)
	}

	// 如果传了最小年龄 (Age >= ?)
	if params.MinAge != nil {
		query = query.Where("age >= ?", *params.MinAge)
	}

	// 如果传了最大年龄 (Age <= ?)
	if params.MaxAge != nil {
		query = query.Where("age <= ?", *params.MaxAge)
	}

	// 3. 最后发射！执行查询
	// 只有这一步才会真正去数据库拿数据
	result := query.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
