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

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	version := beego.AppConfig.String("version")
	topic := new(models.Topic)
	topics, _ := topic.GetTopics(0, 10)
	this.Data["Topics"] = topics
	this.Data["Version"] = version
	this.TplName = "index.html"

}
