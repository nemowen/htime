package main

import (
	"github.com/astaxie/beego"
	_ "github.com/nemowen/htime/models"
	_ "github.com/nemowen/htime/routers"
)

func main() {
	beego.Run()
}
