package user_models

import (
	"im_server/common/ctype"
	"im_server/common/models"
)

// UserConfModel 用户配置表
type UserConfModel struct {
	models.Models
	UserID               uint                        `json:"userID" gorm:"comment:'用户ID'"`
	UserModel            UserModel                   `gorm:"foreignKey:UserID" json:"-"` // 用户模型外键
	RecallMessage        *string                     `gorm:"size:32;comment:'撤回消息的提示内容'" json:"recallMessage"`
	FriendOnline         bool                        `gorm:"default:true;comment:'好友上线提醒'" json:"friendOnline"`
	EnableSound          bool                        `gorm:"default:true;comment:'是否开启声音提醒'" json:"enableSound"`
	SecureLink           bool                        `gorm:"default:false;comment:'安全链接是否开启'" json:"secureLink"`
	SavePwd              bool                        `gorm:"default:false;comment:'是否保存密码'" json:"savePwd"`
	SearchUser           int8                        `gorm:"default:0;comment:'别人查找到你的方式: 0 不允许, 1 用户号, 2 昵称'" json:"searchUser"`
	Verification         int8                        `gorm:"default:2;comment:'验证类型: 0 不允许任何人, 1 允许任何人, 2 验证消息, 3 回答问题, 4 正确回答问题'" json:"verification"`
	VerificationQuestion *ctype.VerificationQuestion `gorm:"type:json;comment:'验证问题'" json:"verificationQuestion"`
	Online               bool                        `gorm:"default:false;comment:'是否在线'" json:"online"`
}
