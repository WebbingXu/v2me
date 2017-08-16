package controllers

import (

	"os"

	"common"

	"github.com/astaxie/beego"
)

type MarkDownController struct {
	beego.Controller
}


func (c *MarkDownController) Home() {
	c.TplName = "home.html"
}

func (c *MarkDownController) Edit() {
	c.TplName = "editor_page.html"
}

func (c *MarkDownController) GetEditContent() {
	Content := c.GetString("Content")
	FileName := c.GetString("FileName")

	FileContent := common.ArticleTemplateFront + Content + common.ArticleTemplateRear
	f, err := os.Create("./articles/" + FileName + ".html")
	if err != nil{
		c.Ctx.WriteString("create file failed: " + err.Error())
		return
	}
	f.WriteString(FileContent)
	f.Close()
	c.Ctx.WriteString("ok")
}