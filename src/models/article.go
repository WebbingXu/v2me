package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct{
	ArticleName string
	CreatedTime string

}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Article))
}
