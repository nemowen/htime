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
	"github.com/astaxie/beego/validation"
	"github.com/nemowen/htime/models"
	"github.com/nemowen/htime/utils"
	"strings"
)

type UserController struct {
	baseController
}

//用户列表
func (this *UserController) List() {
	var page int64
	var pagesize int64 = 10
	var list []*models.User
	var user models.User

	if page, _ = this.GetInt64("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize

	count, _ := user.AllUserCount()
	if count > 0 {
		list, _ = user.GetUsers(int(offset), int(pagesize))
	}

	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page, count, pagesize, "/admin/user/list?page=%d").ToString()
	this.display("user_list")
}

//添加用户
func (this *UserController) Add() {
	input := make(map[string]string)
	errmsg := make(map[string]string)
	if this.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		password2 := strings.TrimSpace(this.GetString("password2"))
		email := strings.TrimSpace(this.GetString("email"))
		active, _ := this.GetInt64("active")

		input["username"] = username
		input["password"] = password
		input["password2"] = password2
		input["email"] = email

		valid := validation.Validation{}

		if v := valid.Required(username, "username"); !v.Ok {
			errmsg["username"] = "请输入用户名"
		} else if v := valid.MaxSize(username, 15, "username"); !v.Ok {
			errmsg["username"] = "用户名长度不能大于15个字符"
		}

		if v := valid.Required(password, "password"); !v.Ok {
			errmsg["password"] = "请输入密码"
		}

		if v := valid.Required(password2, "password2"); !v.Ok {
			errmsg["password2"] = "请再次输入密码"
		} else if password != password2 {
			errmsg["password2"] = "两次输入的密码不一致"
		}

		if v := valid.Required(email, "email"); !v.Ok {
			errmsg["email"] = "请输入email地址"
		} else if v := valid.Email(email, "email"); !v.Ok {
			errmsg["email"] = "Email无效"
		}

		if active > 0 {
			active = 1
		} else {
			active = 0
		}

		if len(errmsg) == 0 {
			var user models.User
			user.Username = username
			user.Password = utils.EncodeByMd5(password)
			user.Email = email
			user.Active = int8(active)
			if err := user.SaveUser(); err != nil {
				this.showMsg(err.Error())
			}
			this.Redirect("/admin/user/list", 302)
		}

	}

	this.Data["input"] = input
	this.Data["errmsg"] = errmsg
	this.display("user_add")
}

//编辑用户
func (this *UserController) Edit() {
	id, _ := this.GetInt64("id")
	user := new(models.User)
	if err := user.GetUserById(id); err != nil {
		this.showMsg("用户不存在")
	}

	errmsg := make(map[string]string)

	if this.Ctx.Request.Method == "POST" {
		password := strings.TrimSpace(this.GetString("password"))
		password2 := strings.TrimSpace(this.GetString("password2"))
		email := strings.TrimSpace(this.GetString("email"))
		active, _ := this.GetInt64("active")
		valid := validation.Validation{}

		if password != "" {
			if v := valid.Required(password2, "password2"); !v.Ok {
				errmsg["password2"] = "请再次输入密码"
			} else if password != password2 {
				errmsg["password2"] = "两次输入的密码不一致"
			} else {
				user.Password = utils.EncodeByMd5(password)
			}
		}
		if v := valid.Required(email, "email"); !v.Ok {
			errmsg["email"] = "请输入email地址"
		} else if v := valid.Email(email, "email"); !v.Ok {
			errmsg["email"] = "Email无效"
		} else {
			user.Email = email
		}

		if active > 0 {
			user.Active = 1
		} else {
			user.Active = 0
		}

		if len(errmsg) == 0 {
			user.UpdateUser()
			this.Redirect("/admin/user/list", 302)
		}
	}
	this.Data["errmsg"] = errmsg
	this.Data["user"] = user
	this.display("user_edit")
}

//删除用户
func (this *UserController) Delete() {
	id, _ := this.GetInt64("id")
	if id == 1 {
		this.showMsg("不能删除ID为1的用户")
	}
	user := new(models.User)
	if user.GetUserById(id) == nil {
		user.Delete()
	}

	this.Redirect("/admin/user/list", 302)
}
