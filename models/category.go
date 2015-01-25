package models

import (
	"time"
)

type Category struct {
	Id          int64     `xorm:pk autoincr`
	Name        string    `xorm:"varchar(100) notnull"` // 分类名称
	TopicNumber int64     `xorm:"int"`                  // 文章数量
	CreateTime  time.Time `xorm:DateTime created`       // 创建时间
}
