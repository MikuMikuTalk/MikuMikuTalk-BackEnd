package user_models

import "im_server/common/models"

type UserModel struct {
	models.Models
	Pwd            string         `gorm:"size:64;comment:'密码'" json:"pwd"`                   // 密码
	Nickname       string         `gorm:"size:32;comment:'用户名';unique" json:"nickname"`      // 用户名 unique防止用户重名
	Abstract       string         `gorm:"size:128;comment:'简介'" json:"abstract"`             // 简介
	Avatar         string         `gorm:"size:256;comment:'头像'" json:"avatar"`               // 头像
	IP             string         `gorm:"size:32;comment:'ip地址'" json:"ip"`                  // ip地址
	Addr           string         `gorm:"size:64;comment:'地址'" json:"addr"`                  // 地址
	Role           int8           `json:"role" gorm:"comment:'角色 1是管理员 2是普通用户'"`             // 角色 1 管理员  2 普通用户
	RegisterSource string         `gorm:"size:16;comment:'注册来源'" json:"registerSource"`      // 注册来源
	UserConfModel  *UserConfModel `gorm:"foreignKey:UserID;comment:''" json:"UserConfModel"` // 用户配置外键
}
