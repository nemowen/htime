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
	"strconv"
	"strings"
)

type LoginController struct {
	baseController
}

func (l *LoginController) Login() {
	if l.GetString("login") == "yes" {
		account := strings.TrimSpace(l.GetString("account"))
		password := strings.TrimSpace(l.GetString("password"))
		remember := l.GetString("remember")
		if account != "" && password != "" {

			// find user in database with the account
			dbuser := new(models.User)
			err := dbuser.GetUserByUsername(account)
			if err != nil || dbuser == nil || dbuser.Password != utils.EncodeByMd5(password) {
				l.Data["errmsg"] = "用户名或密码有误"
			} else if dbuser.Active == 0 {
				l.Data["errmsg"] = "该帐号未激活"
			} else {
				dbuser.LoginIp = l.Ctx.Input.IP()
				dbuser.LoginCount += 1
				err := dbuser.UpdateUser()
				if err != nil {
					println(err.Error())
				}
				// remember me
				authkey := utils.EncodeByMd5(dbuser.LoginIp + "|" + dbuser.Password)
				if remember == "yes" {
					l.Ctx.SetCookie("auth", strconv.FormatInt(dbuser.Id, 10)+"|"+authkey, 7*86400)
				} else {
					l.Ctx.SetCookie("auth", strconv.FormatInt(dbuser.Id, 10)+"|"+authkey)
				}

				l.Redirect("/admin", 302)
			}
		}
	}
	l.TplNames = "admin/login.html"
}

//退出登录
func (this *LoginController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect("/admin/login", 302)
}
