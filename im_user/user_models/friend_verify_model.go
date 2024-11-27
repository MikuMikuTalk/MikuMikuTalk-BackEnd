package user_models

import (
	"im_server/common/ctype"
	"im_server/common/models"
)

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Models
	SendUserID           uint                        `gorm:"index;comment:'发起验证方用户ID'" json:"sendUserID"`
	SendUserModel        UserModel                   `gorm:"foreignKey:SendUserID" json:"-"` // 发起验证方用户模型
	RevUserID            uint                        `gorm:"index;comment:'接收验证方用户ID'" json:"revUserID"`
	RevUserModel         UserModel                   `gorm:"foreignKey:RevUserID" json:"-"` // 接收验证方用户模型
	Status               int8                        `gorm:"default:0;comment:'状态: 0 未操作, 1 同意, 2 拒绝, 3 忽略 4 删除'" json:"status"`
	AdditionalMessages   string                      `gorm:"size:128;comment:'附加消息'" json:"additionalMessages"`
	VerificationQuestion *ctype.VerificationQuestion `gorm:"type:json;comment:'验证问题, 仅验证方式为3或4时需要'" json:"verificationQuestion"`
}
