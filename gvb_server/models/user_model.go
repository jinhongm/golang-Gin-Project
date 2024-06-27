package models

import "gvb_server/models/ctype"

// GORM标签用于定义数据库的细节（如primarykey和字段大小），
// 而JSON标签定义了如何将这些字段转换为JSON（例如，当从API返回响应时字段的名称
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:42" json:"nick_name"`
	UserName   string           `gorm:"size:36" json:"user_name"`
	PassWord   string           `gorm:"size:128" json:"password"`
	Avatar     string           `gorm:"size:256" json:"-"` // Avatar
	Email      string           `gorm:"size:128" json:"email"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr"`
	Token      string           `gorm:"size:64" json:"token"`
	IP         string           `gorm:"size:20" json:"ip"`
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role"`
	SignStatus ctype.SignStatus `gorm:"type:smallint(6)" json:"sign_status"`
}
