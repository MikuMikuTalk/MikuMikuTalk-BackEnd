package user_models

import "im_server/common/models"

type UserModel struct {
	models.Models
	Pwd            string `gorm:"size:64" json:"pwd"`            // 密码
	Nickname       string `gorm:"size:32" json:"nickname"`       // 用户名
	Abstract       string `gorm:"size:128" json:"abstract"`      // 简介
	Avatar         string `gorm:"size:256" json:"avatar"`        // 头像
	IP             string `gorm:"size:32" json:"ip"`             // ip地址
	Addr           string `gorm:"size:64" json:"addr"`           // 地址
	Role           int8   `json:"role"`                          // 角色 1 管理员  2 普通用户
	RegisterSource string `gorm:"size:16" json:"registerSource"` // 注册来源
}
