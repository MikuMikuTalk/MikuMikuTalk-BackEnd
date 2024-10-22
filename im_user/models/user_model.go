package models

import "im_server/common/models"

type UserModel struct {
	models.Models
	Pwd      string `gorm:"size:64" json:"pwd"`
	Nickname string `gorm:"size:32" json:"nickname"`
	Abstract string `gorm:"size:128" json:"abstract"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	IP       string `gorm:"size:32" json:"ip"`
	Addr     string `gorm:"size:64" json:"addr"`
}
