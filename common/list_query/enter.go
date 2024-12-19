package list_query

import (
	"fmt"

	"im_server/common/models"

	"gorm.io/gorm"
)

type Option struct {
	PageInfo models.PageInfo
	Where    *gorm.DB // 高级查询
	Debug    bool     // 调试
	Joins    string
	Likes    []string             // 模糊匹配的字段
	Preload  []string             // 预加载字段
	Table    func() (string, any) // 子查询
	Groups   []string             // 分组
}

func ListQuery[T any](db *gorm.DB, model T, option Option) (list []T, count int64, err error) {
	if option.Debug {
		db = db.Debug()
	}

	// 把结构体自己的查询条件查了
	query := db.Where(model)

	// 模糊匹配
	if option.PageInfo.Key != "" && len(option.Likes) > 0 {
		likeQuery := db.Where("")
		for index, column := range option.Likes {
			if index == 0 {
				// where name like '%fengfeng%'
				likeQuery.Where(fmt.Sprintf("%s like '%%?%%'", column), option.PageInfo.Key)
			} else {
				likeQuery.Or(fmt.Sprintf("%s like '%%?%%'", column), option.PageInfo.Key)
			}
		}
		query.Where(likeQuery)
	}

	if option.Table != nil {
		table, data := option.Table()
		query = query.Table(table, data)
	}

	if option.Joins != "" {
		query = query.Joins(option.Joins)
	}

	if option.Where != nil {
		query = query.Where(option.Where)
	}

	if len(option.Groups) > 0 {
		for _, group := range option.Groups {
			query = query.Group(group)
		}
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
	if option.PageInfo.Limit != -1 { // 如果是-1就是查全部
		if option.PageInfo.Limit <= 0 {
			option.PageInfo.Limit = 10
		}
	}
	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit

	if option.PageInfo.Sort != "" {
		query.Order(option.PageInfo.Sort)
	}

	err = query.Limit(option.PageInfo.Limit).Offset(offset).Find(&list).Error
	return
}
