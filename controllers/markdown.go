package controllers

import (
	"github.com/astaxie/beego"
	"os"
)

type MarkDownController struct {
	beego.Controller
}


func (c *MarkDownController) Home() {
	c.TplName = "index.html"
}

func (c *MarkDownController) Edit() {
	c.TplName = "editor.md/examples/simple.html"
}

func (c *MarkDownController) GetEditContent() {
	Content := c.GetString("Content")
	if Content == "" {
		c.Ctx.WriteString("please input content")
		return
	}
	c.Data["Content"] = Content






	tmplete := `
		<!DOCTYPE html>
		<html lang="zh">
		<head>
		<meta charset="utf-8" />
		<link rel="stylesheet" href="../mdeditor/examples/css/style.css" />
		<link rel="stylesheet" href="../mdeditor/css/editormd.css" />
		<script src="../mdeditor/examples/js/jquery.min.js"></script>
		<script src="../mdeditor/editormd.min.js"></script>
		</head>

		<body >
		<div style="border: 1px solid #000000; width: 300px; height: 500px;margin: 0 auto;">
		`
	tmplete += Content
	tmplete += `
		</div>
		</body>
		</html>
		`
	f, err := os.Create("./articles/a.html")
	if err != nil{
		c.Ctx.WriteString("create file failed: " + err.Error())
		return
	}
	f.WriteString(tmplete)
	f.Close()
	c.Ctx.WriteString("ok")
}