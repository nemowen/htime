package models

import (
	"time"
)

type Topic struct {
	Id          int64     `xorm:"pk autoincr"`
	Title       string    `xorm:"varchar(255) notnull"` // 标题
	Text        string    `xorm:"text"`                 // 内容
	Images      string    `xorm:"varchar(255)"`         // 图片
	Flags       string    `xorm:"varchar(255)"`         // 标签:多个标签使用|分隔
	CategorieId int64     `xorm:index`                  // 分类ID
	Categorie   *Category `xorm:"- <- ->"`              // 分类
	AuthorId    int64     `xorm:index`                  // 作者ID
	Author      *User     `xorm:"- <- ->"`              // 作者
	SourceFrom  string    `xorm:"varchar(255)"`         // 信息来源
	CreateTime  time.Time `xorm:"DateTime created"`     // 创建时间
	UpdateTime  time.Time `xorm:"DateTime updated"`     // 更新时间
	Version     int32     `xorm:"version"`
}
