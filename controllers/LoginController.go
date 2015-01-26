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
	"github.com/astaxie/beego"
	"github.com/nemowen/htime/models"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.Data["LoginActive"] = true
	l.TplNames = "admin/login.html"
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")

	l.Data["LoginActive"] = true
	l.TplNames = "admin/login.html"
	if len(username) == 0 || len(password) == 0 {
		l.Data["ErrorMessage"] = "用户名与密码不能为空"
	}

	// 查询数据库中的用户
	dbuser, err := models.GetUserByUsername(username)
	if err != nil || dbuser == nil || dbuser.Password != password {
		l.Data["ErrorMessage"] = "用户名或密码有误"
		return
	}

	if dbuser.Password == password {
		l.TplNames = "admin/login.html"
	}
}

func (l *LoginController) SignupPage() {
	l.Data["SignupActive"] = true
	l.TplNames = "admin/login.html"
}

func (l *LoginController) Signup() {
	username := l.GetString("username")
	password := l.GetString("password")
	repassword := l.GetString("repassword")
	email := l.GetString("email")

	l.TplNames = "admin/login.html"
	if len(username) == 0 || len(password) == 0 || len(repassword) == 0 || len(email) == 0 {
		l.Data["ErrorMessage"] = "用户名,密码,email,不能为空"
		l.Data["SignupActive"] = true
		return
	}

	user := new(models.User)
	user.Username = username
	user.Password = password
	user.Email = email

	_, err := models.CreateUser(user)
	if err != nil {
		l.Data["SignupActive"] = true
		l.Data["ErrorMessage"] = err.Error()
		return
	}

	l.Data["LoginActive"] = true
	l.TplNames = "admin/login.html"

	l.Data["ErrorMessage"] = "success"

}
