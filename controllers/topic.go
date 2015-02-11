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
	"fmt"
	"github.com/nemowen/htime/models"
	"strconv"
	"strings"
)

type TopicController struct {
	baseController
}

// 文章列表
func (this *TopicController) List() {
	var (
		page       int64
		pagesize   int64 = 10
		status     int64
		offset     int64
		topic      models.Topic
		searchtype string
		keyword    string
	)
	list := make([]*models.Topic, 0, pagesize)
	searchtype = this.GetString("searchtype")
	keyword = this.GetString("keyword")
	status, _ = this.GetInt64("status")
	if page, _ = this.GetInt64("page"); page < 1 {
		page = 1
	}

	offset = (page - 1) * pagesize
	topic.Status = int8(status)

	if keyword != "" {
		switch searchtype {
		case "title":
			topic.Title = keyword
		case "tag":
			topic.Tags = keyword
		}
	}

	count, _ := topic.Count()

	if count > 0 {
		list, _ = topic.GetTopics(int(offset), int(pagesize))
	}

	this.Data["searchtype"] = searchtype
	this.Data["keyword"] = keyword
	topic.Status = 0
	this.Data["count_0"], _ = topic.Count()
	topic.Status = 1
	this.Data["count_1"], _ = topic.Count()
	topic.Status = 2
	this.Data["count_2"], _ = topic.Count()
	this.Data["status"] = status
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page, count, pagesize, fmt.Sprintf(
		"/admin/topic/list?status=%d&searchtype=%s&keyword=%s&page=%s", status,
		searchtype, keyword, "%d")).ToString()

	this.display("topic_list")
}

func (this *TopicController) Add() {
	this.display("topic_add")
}

func (this *TopicController) Save() {
	topic := new(models.Topic)
	id, _ := this.GetInt64("id")
	topic.Title = strings.TrimSpace(this.GetString("title"))
	topic.SourceFrom = strings.TrimSpace(this.GetString("sourcefrom"))
	topic.Content = this.GetString("content")
	topic.Tags = strings.TrimSpace(this.GetString("tags"))
	topic.TitleColor = strings.TrimSpace(this.GetString("color"))
	topic.Cover = strings.TrimSpace(this.GetString("cover"))
	topic.Status = 0
	topic.CategorieId = 0
	topic.IsTop = 0

	if topic.Title == "" {
		this.showmsg("标题不能为空！")
	}

	status, _ := this.GetInt64("status")

	if this.GetString("istop") == "1" {
		topic.IsTop = 1
	}

	if status != 1 && status != 2 {
		status = 0
	}
	topic.Status = int8(status)
	if topic.Cover == "" {
		topic.Cover = "/static/upload/defaultcover.png"
	}

	topic.AuthorId = this.userid
	if id > 0 {
		topic.Id = id
		topic.Update()
	} else {
		topic.Save()
	}
	this.Redirect("/admin/topic/list", 302)
}

// 删除
func (this *TopicController) Delete() {
	id, _ := this.GetInt64("id")
	topic := models.Topic{}
	topic.DeleteById(id)
	this.Redirect("/admin/topic/list", 302)
}

// 编辑
func (this *TopicController) Edit() {
	id, _ := this.GetInt64("id")
	topic := models.Topic{}
	if topic.GetTopicById(id) != nil {
		this.Abort("404")
	}
	this.Data["topic"] = topic
	this.display("topic_edit")
}

// 批处理
func (this *TopicController) Batch() {
	ids := this.GetStrings("ids[]")
	op := this.GetString("op")

	idarr := make([]int64, 0, 10)
	for _, v := range ids {
		if id, _ := strconv.Atoi(v); id > 0 {
			idarr = append(idarr, int64(id))
		}
	}

	var topic models.Topic
	switch op {
	case "topub": //移到已发布
		topic.Status = 0
		models.GetSession().In("id", idarr).Cols("status").Update(topic)
	case "todrafts": //移到草稿箱
		topic.Status = 1
		models.GetSession().In("id", idarr).Cols("status").Update(topic)
	case "totrash": //移到回收站
		topic.Status = 2
		models.GetSession().In("id", idarr).Cols("status").Update(topic)
	case "delete": //批量删除
		for _, id := range idarr {
			topic.DeleteById(id)
		}
	}
	this.Redirect(this.Ctx.Request.Referer(), 302)
}
