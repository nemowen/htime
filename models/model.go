package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 数据库初始化
func init() {
	orm.RegisterModel(new(Topic), new(User))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:wenbin@/htime?charset=utf8", 10, 100)
	orm.Debug = true
	orm.RunSyncdb("default", true, true)
}

type Topic struct {
	Id         int64
	Title      string    // 标题
	Text       string    `orm:"size(5000)"` //内容
	Images     string    // 图片
	Flags      string    // 标签:多个标签使用|分隔
	Categories *Category `orm:"rel(fk)"` // 分类
	Author     *User     `orm:"rel(fk)"` // 作者
	SourceFrom string    `orm:"null"`    // 信息来源
	CreateTime time.Time `orm:"index"`   // 创建时间
}

type User struct {
	Id            int64
	Name          string    // `姓名`
	Email         string    // `邮箱地址`
	Username      string    // `用户名称`
	Password      string    // `用户密码`
	CreateTime    time.Time // `创建时间`
	LastLoginTime time.Time // `最后登录时间`
	LoginIp       string    // `登录IP`
}

type Category struct {
	Id          int64
	Name        string //分类名称
	TopicNumber int64  // 文章数量
}
