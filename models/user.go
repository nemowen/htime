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

type User struct {
	Id            int64     `xorm:"pk autoincr"`
	Name          string    `xorm:varchar(100)`         // `姓名`
	Email         string    `xorm:varchar(100)`         // `邮箱地址`
	Username      string    `xorm:varchar(100) notnull` // `用户名称`
	Password      string    `xorm:varchar(100) notnull` // `用户密码`
	LoginCount    int64     `xorm:int`                  // 登录次数
	Authkey       string    `xorm:varchar(10)"`
	Active        int8      `xorm:int`                // 是否激活
	CreateTime    time.Time `xorm:"DateTime created"` // `创建时间`
	LastLoginTime time.Time `xorm:"DateTime updated"` // `最后登录时间`
	LoginIp       string    `xorm:"varchar(100)"`     // `登录IP`
}

var (
	ErrUserNotExist    = errors.New("User does not exist")
	ErrUserAlreadExist = errors.New("User alread exist")
	ErrUserIsNull      = errors.New("User is null")
)

// Find user with id
func (u *User) GetUserById(id int64) error {
	// id conditon
	b, err := orm.Where("id=?", id).Get(u)
	if err != nil {
		return err
	} else if !b {
		return ErrUserNotExist
	}
	return nil
}

// Find user with username
func (u *User) GetUserByUsername(username string) error {
	if len(username) == 0 {
		return ErrParameter
	}

	b, err := orm.Where("username=?", username).Get(u)
	if err != nil {
		return err
	} else if !b {
		return ErrUserNotExist
	}
	return nil
}

// Find user with e-mail address
func (u *User) GetUserByEmail(email string) error {
	if len(email) == 0 {
		return ErrParameter
	}

	b, err := orm.Where("email=?", email).Get(u)
	if err != nil {
		return err
	} else if !b {
		return ErrUserNotExist
	}
	return nil
}

// Check user whether or not exist in db with the username
func (u *User) IsUserExist() (bool, error) {
	if len(u.Username) == 0 {
		return false, ErrParameter
	}
	return orm.Get(u)
}

// Save user entity
func (u *User) SaveUser() error {
	if u == nil {
		return ErrUserIsNull
	}
	u.fixData()
	b, err := u.IsUserExist()
	if err != nil {
		return err
	} else if b {
		return ErrUserAlreadExist
	}
	_, err = orm.InsertOne(u)
	return err
}

// Update user entity
func (u *User) UpdateUser() error {
	u.fixData()
	_, err := orm.Id(u.Id).AllCols().Update(u)
	return err
}

func (u *User) AllUserCount() (int64, error) {
	total, err := orm.Count(u)
	return total, err
}

func (u *User) DisableUserCount() (int64, error) {
	total, err := orm.Where("Active = ?", 0).Count(u)
	return total, err
}

// GetUsers method returns the user list
func (u *User) GetUsers(offset int, size int) ([]*User, error) {
	if size == 0 {
		size = MIN_PAGE_SIZE
	}
	if size > MAX_PAGE_SIZE {
		size = MAX_PAGE_SIZE
	}
	users := make([]*User, 0, size)
	err := orm.Limit(size, offset).Desc("id").Find(&users)
	return users, err
}

func (u *User) Delete() error {
	_, err := orm.Delete(u)
	return err
}

func (u *User) fixData() {
	if len(u.Name) > 100 {
		u.Name = u.Name[:100]
	}
	if len(u.Email) > 100 {
		u.Email = u.Email[:100]
	}
	if len(u.Username) > 100 {
		u.Username = u.Username[:100]
	}
}

func (u *User) String() string {
	return u.Username
}
