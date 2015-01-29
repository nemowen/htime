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

var (
	ErrTopicNotExist = errors.New("Topic does not exist")
)

// Save topic to db
func CreateTopic(t *Topic) error {
	if len(t.Title) > 255 {
		t.Title = t.Title[:255]
	}
	if len(t.Images) > 255 {
		t.Images = t.Images[:255]
	}
	if len(t.Flags) > 255 {
		t.Flags = t.Flags[:255]
	}
	if len(t.SourceFrom) > 255 {
		t.SourceFrom = t.SourceFrom[:255]
	}
	_, err := orm.InsertOne(t)
	return err
}

// Delete topic with id
func DeleteTopicById(id int64) error {
	topic, err := GetTopicById(id)
	if err != nil {
		return nil, err
	}
	err = orm.Delete(topic)
	return err
}

// Find a topic with id
func GetTopicById(id int64) (*Topic, error) {
	topic := new(Topic)

	// id conditon
	b, err := orm.Where("id=?", id).Get(topic)
	if err != nil {
		return nil, err
	} else if !b {
		return nil, ErrTopicNotExist
	}
	return topic, nil
}

// GetTopics method returns the topic list
func GetTopics(offset int64, size int64) ([]*Topic, error) {
	if size == 0 {
		size = 20
	}
	if size > 50 {
		size = 50
	}

	topics := make([]*Topic, 0, size)
	session := orm.Limit(size, offset).Desc("id")
	err := session.Find(topics)
	return topics, err
}
