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
package controllers

import (
	"github.com/nemowen/htime/models"
	"github.com/nemowen/htime/utils"
	"os"
	"runtime"
	"strings"
)

type MainController struct {
	baseController
}

func (this *MainController) Index() {
	this.Data["hostname"], _ = os.Hostname()
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH
	//this.Data["postnum"], _ = new(models.Post).Query().Count()
	//this.Data["tagnum"], _ = new(models.Tag).Query().Count()
	user := new(models.User)
	this.Data["DisableUserNum"], _ = user.DisableUserCount()
	this.Data["AllUserNum"], _ = user.AllUserCount()
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/index.html"
}

//资料修改
func (this *MainController) Profile() {
	user := new(models.User)
	if err := user.GetUserById(this.userid); err != nil {
		this.showMsg(err.Error())
	}
	if this.Ctx.Request.Method == "POST" {
		errmsg := make(map[string]string)
		password := strings.TrimSpace(this.GetString("password"))
		newpassword := strings.TrimSpace(this.GetString("newpassword"))
		newpassword2 := strings.TrimSpace(this.GetString("newpassword2"))
		updated := false
		if newpassword != "" {
			if password == "" || utils.EncodeByMd5(password) != user.Password {
				errmsg["password"] = "当前密码错误"
			} else if len(newpassword) < 6 {
				errmsg["newpassword"] = "密码长度不能少于6个字符"
			} else if newpassword != newpassword2 {
				errmsg["newpassword2"] = "两次输入的密码不一致"
			}
			if len(errmsg) == 0 {
				user.Password = utils.EncodeByMd5(newpassword)
				user.UpdateUser()
				updated = true
			}
		}
		this.Data["updated"] = updated
		this.Data["errmsg"] = errmsg
	}
	this.Data["user"] = user
	this.display()
}
