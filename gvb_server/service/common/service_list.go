package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB

	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	if option.Sort == "" {
		option.Sort = "created_at asc" // 默认按时间降序排序
	}

	err = DB.Model(&model).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (option.Page - 1) * option.Limit
	if offset < 0 || option.Page == 0 {
		offset = 0
	}
	query := DB.Where(model).Order(option.Sort)
	if option.Limit > 0 {
		query = query.Limit(option.Limit)
	}
	query = query.Offset(offset)

	err = query.Find(&list).Error
	return list, count, err
}
