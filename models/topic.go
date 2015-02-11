// Copyright 2014 nemowen
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// An open source project for personal blog website
package models

import (
	"errors"
	"time"
)

type Topic struct {
	Id          int64     `xorm:"pk autoincr"`
	Title       string    `xorm:"varchar(255) notnull"` // 标题
	TitleColor  string    `xorm:varchar(7)`             // 标题置顶颜色
	Content     string    `xorm:"text"`                 // 内容
	Tags        string    `xorm:"varchar(255)"`         // 标签:多个标签使用|分隔
	CategorieId int64     `xorm:index`                  // 分类ID
	Categorie   *Category `xorm:"- <- ->"`              // 分类
	AuthorId    int64     `xorm:index`                  // 作者ID
	Author      *User     `xorm:"- <- ->"`              // 作者
	Status      int8      `xorm:int`                    // 发布状态
	IsTop       int8      `xorm:int`                    // 是否致顶
	Cover       string    `xorm:varchar(100)`           // 封面图片
	Views       int64     `xorm:int`                    // 阅读数
	SourceFrom  string    `xorm:"varchar(255)"`         // 信息来源
	CreateTime  time.Time `xorm:"DateTime created"`     // 创建时间
	UpdateTime  time.Time `xorm:"DateTime updated"`     // 更新时间
}

var (
	ErrTopicNotExist = errors.New("Topic does not exist")
	ErrTopicIsNull   = errors.New("Topic is null")
)

func (t *Topic) fixData() {
	if len(t.Title) > 255 {
		t.Title = t.Title[:255]
	}
	if len(t.Tags) > 255 {
		t.Tags = t.Tags[:255]
	}
	if len(t.SourceFrom) > 255 {
		t.SourceFrom = t.SourceFrom[:255]
	}
}

// Save topic to db
func (t *Topic) Save() error {
	if t == nil {
		return ErrTopicIsNull
	}
	t.fixData()
	_, err := orm.InsertOne(t)
	return err
}

// Delete topic with id
func (t *Topic) DeleteById(id int64) error {
	err := t.GetTopicById(id)
	if err != nil {
		return err
	}
	_, err = orm.Delete(t)
	return err
}

// update topic
func (t *Topic) Update() error {
	t.fixData()
	_, err := orm.Id(t.Id).AllCols().Update(t)
	return err
}

// Find a topic with id
func (t *Topic) GetTopicById(id int64) error {
	// id conditon
	b, err := orm.Where("id=?", id).Get(t)
	if err != nil {
		return err
	} else if !b {
		return ErrTopicNotExist
	}
	return nil
}

// GetTopics method returns the topic list
func (t *Topic) GetTopics(offset int, size int) ([]*Topic, error) {
	if size == 0 {
		size = 10
	}
	if size > 50 {
		size = 50
	}
	topics := make([]*Topic, 0, size)
	session := orm.Where("status=?", t.Status)
	if len(t.Title) > 0 {
		session.And("title=?", t.Title)
	} else if len(t.Tags) > 0 {
		session.And("tags=?", t.Tags)
	}
	err := session.Limit(size, offset).Desc("id").Find(&topics)
	return topics, err
}

func (t *Topic) Count() (int64, error) {
	session := orm.Where("status=?", t.Status)
	if len(t.Title) > 0 {
		session.And("title=?", t.Title)
	} else if len(t.Tags) > 0 {
		session.And("tags=?", t.Tags)
	}

	return session.Count(t)
}
