package main

import (
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/nemowen/htime/models"
	_ "github.com/nemowen/htime/routers"
)

func Split(s string, d string) []string {
	arr := strings.Split(s, d)
	return arr
}

func main() {
	beego.AddFuncMap("Split", Split)
	beego.Run()
}
