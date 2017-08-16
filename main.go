package main

import (
	_ "routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}


func main() {
	beego.SetStaticPath("/markdown", "markdown")
	beego.SetStaticPath("/articles", "articles")
	beego.Run()
}

