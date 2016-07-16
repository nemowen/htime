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
	"strconv"
	"strings"
)

const (
	BIG_PIC_PATH   = "./static/upload/bigpic/"
	SMALL_PIC_PATH = "./static/upload/smallpic/"
	FILE_PATH      = "./static/upload/attachment/"
)

var pathArr []string = []string{"", BIG_PIC_PATH, SMALL_PIC_PATH, FILE_PATH}

type baseController struct {
	beego.Controller
	userid         int64
	username       string
	moduleName     string
	controllerName string
	actionName     string
}

func (this *baseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.auth()
	this.checkPermission()
}

//登录状态验证
func (this *baseController) auth() {
	if this.controllerName != "login" && this.controllerName != "home" {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr, _ := arr[0], arr[1]
			userid, _ := strconv.ParseInt(idstr, 10, 0)
			if userid > 0 {
				var user models.User
				if user.GetUserById(userid) == nil {
					this.userid = user.Id
					this.username = user.Username
				}
			}
		}
		if this.userid == 0 {
			this.Redirect("/admin/login", 302)
		}
	}
}

//渲染模版
func (this *baseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = "admin/" + tpl[0] + ".html"
	} else {
		tplname = "admin/" + this.actionName + ".html"
	}
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Layout = "admin/layout.html"
	this.TplNames = tplname
}

//权限验证
func (this *baseController) checkPermission() {
	if this.userid != 1 && this.controllerName == "user" {
		this.showmsg("抱歉，只有超级管理员才能进行该操作！")
	}
}

//显示错误提示
func (this *baseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = this.moduleName + "/" + "showmsg.html"
	this.Render()
	this.StopRun()
}
