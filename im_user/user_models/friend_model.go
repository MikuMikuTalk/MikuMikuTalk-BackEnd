package user_models

import "im_server/common/models"

// FriendModel 好友表
type FriendModel struct {
	models.Models
	SendUserID    uint      `json:"sendUserID" gorm:"comment:'发起验证方ID'"`               // 发起验证方
	SendUserModel UserModel `gorm:"foreignKey:SendUserID;comment:'发起验证方的信息'" json:"-"` // 发起验证方
	RevUserID     uint      `json:"revUserID" gorm:"comment:'接收验证方ID'"`                // 接受验证方
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID;comment:'接收验证方信息'" json:"-"`   // 接受验证方
	SenUserNotice string    `gorm:"size:128" json:"senUserNotice"`                     // 发送方备注
	RevUserNotice string    `gorm:"size:128" json:"revUserNotice"`                     // 接收方备注
}

func (f *FriendModel) GetUserNotice(userID uint) string {
	if userID == f.SendUserID {
		//如果我是发起验证方
		return f.SenUserNotice
	}
	if userID == f.RevUserID {
		//如果我是接收方
		return f.RevUserNotice
	}
	return ""
}
