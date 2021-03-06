// Copyright 2017 nemowen
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
	"time"
)

// 相册表
type Photo struct {
	Id       int64     `xorm:int`
	Albumid  int64     `xorm:int`
	Des      string    `xorm:varchar(100)`
	Posttime time.Time `xorm:"DateTime created"`
	Url      string    `xorm:"varchar(70)"`
	Small    string    `xorm:"-"`
}

func (m *Photo) Insert() error {
	if _, err := orm.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *Photo) Read(fields ...string) error {
	if b, err := orm.Cols(fields...).Get(m); !b {
		return err
	}
	return nil
}

func (m *Photo) Update(fields ...string) error {
	if _, err := orm.Cols(fields...).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *Photo) Delete() error {
	if _, err := orm.Delete(m); err != nil {
		return err
	}
	return nil
}

func (t *Photo) GetPhotos(offset int, size int) ([]*Photo, error) {
	if size == 0 {
		size = MIN_PAGE_SIZE
	}
	if size > MAX_PAGE_SIZE {
		size = MAX_PAGE_SIZE
	}
	photos := make([]*Photo, 0, size)
	session := orm.Where("1=1")
	if t.Albumid > 0 {
		session.And("albumid=?", t.Albumid)
	} else if len(t.Des) > 0 {
		session.And("des like '%?%'", t.Des)
	}
	err := session.Limit(size, offset).Desc("posttime,id").Find(&photos)
	return photos, err
}
