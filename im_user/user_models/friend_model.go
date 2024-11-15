package user_models

import "im_server/common/models"

// FriendModel 好友表
type FriendModel struct {
	models.Models
	SendUserID    uint      `json:"sendUserID" gorm:"comment:'发起验证方ID'"`               // 发起验证方
	SendUserModel UserModel `gorm:"foreignKey:SendUserID;comment:'发起验证方的信息'" json:"-"` // 发起验证方
	RevUserID     uint      `json:"revUserID" gorm:"comment:'接收验证方ID'"`                // 接受验证方
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID;comment:'接收验证方信息'" json:"-"`   // 接受验证方
	Notice        string    `gorm:"size:128;comment:'备注'" json:"notice"`               // 备注
}
