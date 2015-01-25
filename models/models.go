package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	orm *xorm.Engine
)

func init() {
	InitDatabase()
}

func InitDatabase() error {
	var err error
	orm, err = xorm.NewEngine("mysql", "root:wenbin@/htime?charset=utf8")
	if err != nil {
		orm.LogError(err.Error())
		return err
	}

	//orm.ShowSQL = true   //则会在控制台打印出生成的SQL语句；
	orm.ShowWarn = true  //则会在控制台打印警告信息；
	orm.ShowDebug = true //则会在控制台打印调试信息；
	orm.ShowErr = true   //则会在控制台打印错误信息；

	// regiest these models
	err = orm.Sync(new(Topic), new(User), new(Category))
	if err != nil {
		orm.LogError(err.Error())
		return err
	}

	orm.LogInfo("init database success.")
	return nil
}
