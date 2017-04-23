// Copyright 2017 nemowen
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
	"strconv"
	"strings"
	"time"

	"github.com/nemowen/htime/models"
)

type AlbumController struct {
	baseController
}

// 相册列表
func (this *AlbumController) List() {
	var page int
	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * models.MIN_PAGE_SIZE

	var album models.Album
	count, _ := album.Count()
	list, _ := album.GetAlbums(offset, models.MIN_PAGE_SIZE)
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(int64(page), count, int64(models.MIN_PAGE_SIZE), "/admin/album/list?page=%d").ToString()
	this.display("album_list")
}

// 创建相册
func (this *AlbumController) Add() {
	if this.Ctx.Request.Method == "POST" {
		rank, _ := this.GetInt("rank")
		var album models.Album
		album.Name = strings.TrimSpace(this.GetString("albumname"))
		album.Cover = strings.TrimSpace(this.GetString("cover"))
		album.Rank = int8(rank)
		album.Posttime = time.Now()
		if err := album.Insert(); err != nil {
			this.showMsg(err.Error())
		}
		this.Redirect("/admin/album/list", 302)

	}
	this.display("album_add")
}

// 删除相册
func (this *AlbumController) Delete() {
	id, _ := this.GetInt64("albumid")
	album := models.Album{Id: id}
	h, _ := strconv.Atoi(this.GetString("ishide"))
	album.Ishide = int8(h)
	if err := album.Update("ishide"); err != nil {
		this.showMsg(err.Error())
		return
	}
	this.Redirect("/admin/album/list", 302)
}

// 修改
func (this *AlbumController) Edit() {
	id, _ := this.GetInt64("albumid")
	album := models.Album{Id: id}
	if album.Read() != nil {
		this.Abort("404")
	}
	if this.Ctx.Request.Method == "POST" {
		rank, _ := this.GetInt("rank")
		album.Cover = this.GetString("cover")
		album.Name = this.GetString("albumname")
		album.Rank = int8(rank)
		album.Update()
		this.Redirect("/admin/album/list", 302)
	}
	this.Data["album"] = album
	this.display("album_edit")
}
