package routers

import (
	"v2me/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "get:Home")
	beego.Router("/edit", &controllers.MainController{}, "get:Edit")
	beego.Router("/get_edit", &controllers.MainController{}, "post:GetEditContent")
}
