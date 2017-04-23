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
type Album struct {
	Id       int64
	Name     string    `xorm:"size(100)"`
	Cover    string    `xorm:"size(70)"`
	Posttime time.Time `xorm:"type(datetime);index"`
	Ishide   int8
	Rank     int8
	Photonum int64
}

func (m *Album) Insert() error {
	if _, err := orm.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *Album) Read(fields ...string) error {
	if b, err := orm.Cols(fields...).Get(m); !b {
		return err
	}
	return nil
}

func (m *Album) Update(fields ...string) error {
	if _, err := orm.Cols(fields...).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *Album) Delete() error {
	if _, err := orm.Delete(m); err != nil {
		return err
	}
	return nil
}

func (t *Album) GetAlbums(offset int, size int) ([]*Album, error) {
	if size == 0 {
		size = MIN_PAGE_SIZE
	}
	if size > MAX_PAGE_SIZE {
		size = MAX_PAGE_SIZE
	}
	albums := make([]*Album, 0, size)
	session := orm.Where("1=1")
	err := session.Limit(size, offset).Desc("id").Find(&albums)
	return albums, err
}

func (t *Album) Count() (int64, error) {
	session := orm.Where("1=1")
	return session.Count(t)
}
