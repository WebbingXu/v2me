package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Home() {
	c.TplName = "index.html"
}

func (c *MainController) Edit() {

	c.TplName = "editor.md/examples/simple.html"
}

func (c *MainController) GetEditContent() {

	content := c.GetString("content")

	if content == "" {
		c.Ctx.WriteString("please input content")
		return
	}
	fmt.Println(content)
	c.Ctx.WriteString("success")

}