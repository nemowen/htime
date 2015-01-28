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
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	orm          *xorm.Engine
	ErrParameter = errors.New("Parameter is wrong")
)

func init() {
	InitDatabase()
}

func InitDatabase() error {
	var err error
	orm, err = xorm.NewEngine("mysql", "root:wenbin@/htime?charset=utf8")
	if err != nil {
		panic(err)
		return err
	}

	//orm.ShowSQL = true   //则会在控制台打印出生成的SQL语句；
	orm.ShowWarn = true  //则会在控制台打印警告信息；
	orm.ShowDebug = true //则会在控制台打印调试信息；
	orm.ShowErr = true   //则会在控制台打印错误信息；

	// regiest these models
	err = orm.Sync(new(Topic), new(User), new(Category))
	if err != nil {
		panic(err)
		return err
	}

	orm.LogInfo("init database success.")
	return nil
}
