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
	CreateTime    time.Time `xorm:"DateTime created"`   // `创建时间`
	LastLoginTime time.Time `xorm:"DateTime updated"`   // `最后登录时间`
	LoginIp       string    `xorm:"varchar(100)"`       // `登录IP`
}

var (
	ErrUserNotExist    = errors.New("User does not exist")
	ErrUserAlreadExist = errors.New("User alread exist")
	ErrUserIsNull      = errors.New("User is null")
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
func CreateUser(u *User) error {
	if u == nil {
		return ErrUserIsNull
	}
	u.fixData()
	b, err := IsUserExist(u.Username)
	if err != nil {
		return err
	} else if b {
		return ErrUserAlreadExist
	}
	_, err = orm.InsertOne(u)
	return err
}

// Update user entity
func UpdateUser(u *User) error {
	if u == nil {
		return ErrUserIsNull
	}
	u.fixData()
	b, err := IsUserExist(u.Username)
	if err != nil {
		return err
	} else if b {
		return ErrUserAlreadExist
	}
	_, err = orm.Id(u.Id).AllCols().Update(u)
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
