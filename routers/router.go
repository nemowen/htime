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
package routers

import (
	"github.com/astaxie/beego"
	"github.com/nemowen/htime/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	// 后台管理
	beego.Router("/admin/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/admin/logout", &controllers.LoginController{}, "*:Logout")

	beego.Router("/admin", &controllers.MainController{}, "*:Index")
	beego.Router("/admin/profile", &controllers.MainController{}, "*:Profile")

	//系统管理
	//beego.Router("/admin/system/setting", &controllers.SystemController{}, "*:Setting")

	// 文章管理
	beego.Router("/admin/topic/add", &controllers.TopicController{}, "*:Add")
	beego.Router("/admin/topic/save", &controllers.TopicController{}, "post:Save")
	beego.Router("/admin/topic/list", &controllers.TopicController{}, "*:List")
	beego.Router("/admin/topic/delete", &controllers.TopicController{}, "*:Delete")
	beego.Router("/admin/topic/edit", &controllers.TopicController{}, "*:Edit")
	beego.Router("/admin/topic/batch", &controllers.TopicController{}, "*:Batch")
	beego.Router("/admin/topic/upload", &controllers.TopicController{}, "*:Upload")

	// 用户管理
	beego.Router("/admin/user/list", &controllers.UserController{}, "*:List")
	beego.Router("/admin/user/add", &controllers.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &controllers.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &controllers.UserController{}, "*:Delete")

	// 照片管理
	beego.Router("/admin/photo/list", &controllers.PhotoController{}, "*:List")
	beego.Router("/admin/photo/cover", &controllers.PhotoController{}, "*:Cover")
	beego.Router("/admin/photo/delete", &controllers.PhotoController{}, "*:Delete")
	beego.Router("/admin/photo/uploadphoto", &controllers.PhotoController{}, "*:UploadPhoto")

	// 相册管理
	beego.Router("/admin/album/add", &controllers.AlbumController{}, "*:Add")
	beego.Router("/admin/album/list", &controllers.AlbumController{}, "*:List")
	beego.Router("/admin/album/edit", &controllers.AlbumController{}, "*:Edit")
	beego.Router("/admin/album/delete", &controllers.AlbumController{}, "*:Delete")

}
