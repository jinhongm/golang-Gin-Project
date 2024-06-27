package models

import "time"

// 记录了用户收藏了什么文章
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}
