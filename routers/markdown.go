package routers

import (
	"github.com/astaxie/beego"
	"v2me/controllers"
)


func init() {

	beego.Router("/", &controllers.MarkDownController{}, "get:Home")
	beego.Router("/edit", &controllers.MarkDownController{}, "get:Edit")
	beego.Router("/get_edit", &controllers.MarkDownController{}, "post:GetEditContent")
}
