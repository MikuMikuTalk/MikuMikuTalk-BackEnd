package group_models

import "im_server/common/models"

// GroupUserTopModel 用户置顶群聊表
type GroupUserTopModel struct {
	models.Models
	UserID  uint `json:"userID"`
	GroupID uint `json:"groupID"`
}
