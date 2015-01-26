// Copyright 2014 HTime authors
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
	CreateTime    time.Time `xorm:"DateTime created"`   // `创建时间`
	LastLoginTime time.Time `xorm:"DateTime updated"`   // `最后登录时间`
	LoginIp       string    `xorm:"varchar(100)"`       // `登录IP`
}

var (
	ErrUserNotExist    = errors.New("User does not exist")
	ErrParameter       = errors.New("Parameter is wrong")
	ErrUserAlreadExist = errors.New("User alread exist")
)

// Find user with id
func GetUserById(id int64) (*User, error) {
	user := new(User)

	// id conditon
	b, err := orm.Where("id=?", id).Get(user)
	if err != nil {
		return nil, err
	} else if !b {
		return nil, ErrUserNotExist
	}
	return user, nil
}

// Find user with username
func GetUserByUsername(username string) (*User, error) {
	if len(username) == 0 {
		return nil, ErrParameter
	}

	user := new(User)
	b, err := orm.Where("username=?", username).Get(user)
	if err != nil {
		return nil, err
	} else if !b {
		return nil, ErrUserNotExist
	}
	return user, nil
}

// Find user with e-mail address
func GetUserByEmail(email string) (*User, error) {
	if len(email) == 0 {
		return nil, ErrParameter
	}

	user := new(User)
	b, err := orm.Where("email=?", email).Get(user)
	if err != nil {
		return nil, err
	} else if !b {
		return nil, ErrUserNotExist
	}
	return user, nil
}

// Check user whether or not exist on db with the username
func IsUserExist(username string) (bool, error) {
	if len(username) == 0 {
		return false, ErrParameter
	}
	return orm.Get(&User{Username: username})
}

// Save user entity
func CreateUser(u *User) (int64, error) {
	b, err := IsUserExist(u.Username)
	if err != nil {
		return -1, err
	} else if b {
		return -1, ErrUserAlreadExist
	}
	u.CreateTime = time.Now()
	return orm.InsertOne(u)
}
