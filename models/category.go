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

type Category struct {
	Id         int64     `xorm:pk autoincr`
	Name       string    `xorm:"varchar(100) notnull"` // 分类名称
	Count      int64     `xorm:"int"`                  // 文章数量
	CreateTime time.Time `xorm:DateTime created`       // 创建时间
}

var (
	ErrCategoryNotExist    = errors.New("Category does not exist")
	ErrCategoryAlreadExist = errors.New("Category alread exist")
	ErrCategoryIsNull      = errors.New("Category is null")
)

func (c *Category) Save() error {
	if len(c.Name) > 100 {
		c.Name = c.Name[:100]
	}
	_, err := orm.InsertOne(c)
	return err
}

func (c *Category) Delete() error {
	b, err := orm.Where("id=?", c.Id).Get(c)
	if err != nil {
		return err
	} else if !b {
		return ErrCategoryNotExist
	}
	_, err = orm.Delete(c)
	return err
}

func (c *Category) Update() error {
	if len(c.Name) > 100 {
		c.Name = c.Name[:100]
	}
	_, err := orm.Id(c.Id).AllCols().Update(c)
	return err
}
