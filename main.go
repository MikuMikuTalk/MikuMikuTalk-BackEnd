package main

import (
	"flag"
	"fmt"
	"im_server/core"
	models2 "im_server/im_chat/chat_models"
	models3 "im_server/im_group/group_models"
	models "im_server/im_user/user_models"
)

type Options struct {
	DB bool
}

func main() {

	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		db := core.InitMysql()
		err := db.AutoMigrate(
			&models.UserModel{},         // 用户表
			&models.FriendModel{},       // 好友表
			&models.FriendVerifyModel{}, // 好友验证表
			&models.UserConfModel{},     // 用户配置表
			&models2.ChatModel{},        // 对话表
			&models3.GroupModel{},       // 群组表
			&models3.GroupMemberModel{}, // 群成员表
			&models3.GroupMsgModel{},    // 群消息表
			&models3.GroupVerifyModel{}, // 群验证表
		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功！")
	}
}
