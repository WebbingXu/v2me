package main

import (
	_ "v2me/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/mdeditor", "mdeditor")

	beego.Run()
}

