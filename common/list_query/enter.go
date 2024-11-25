package list_query

import (
	"fmt"
	"im_server/common/models"

	"gorm.io/gorm"
)

type Option struct {
	PageInfo models.PageInfo
	Where    *gorm.DB //高级查询
	Likes    []string // 模糊匹配得字段
	Preload  []string //预加载字段
}

func ListQuery[T any](db *gorm.DB, model T, option Option) (list []T, count int64, err error) {
	query := db.Where(model)

	if option.PageInfo.Key != "" && len(option.Likes) > 0 {
		likeQuery := db.Where("")
		for index, column := range option.Likes {
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like '%%?%%'", column), option.PageInfo.Key)
			} else {
				likeQuery.Or(fmt.Sprintf("%s like '%%?%%'", column), option.PageInfo.Key)
			}
		}
		query.Where(likeQuery)
	}
	// 求总数
	query.Model(model).Count(&count)

	// 预加载
	for _, s := range option.Preload {
		query = query.Preload(s)
	}

	// 分页查询
	if option.PageInfo.Page <= 0 {
		option.PageInfo.Page = 1
	}
	if option.PageInfo.Limit <= 0 {
		option.PageInfo.Limit = 10
	}

	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit

	err = query.Limit(option.PageInfo.Limit).Offset(offset).Find(&list).Error
	return
}
