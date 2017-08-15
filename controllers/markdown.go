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
	FileName := c.GetString("FileName")

	tmplete := `
		<!DOCTYPE html>
		<html lang="zh">
		<head>
			<meta charset="utf-8" />
			<link rel="stylesheet" href="../mdeditor/examples/css/style.css" />
			<link rel="stylesheet" href="../mdeditor/css/editormd.css" />
			<script src="../mdeditor/examples/js/jquery.min.js"></script>
			<script src="../mdeditor/editormd.min.js"></script>
			<style type="text/css">
				.main_container{
					margin:60px auto;
				width:50%;
				padding: 0px 30px 30px 30px;
				text-align: center;
				line-height: 200%;
				background-color:white;
				}
			</style>
		</head>
		<body style="background-color:rgb(243,243,243)">
		<div class="main_container">
		`
	tmplete += Content
	tmplete += `
		</div>
		</body>
		</html>
		`
	f, err := os.Create("./articles/" + FileName + ".html")
	if err != nil{
		c.Ctx.WriteString("create file failed: " + err.Error())
		return
	}
	f.WriteString(tmplete)
	f.Close()
	c.Ctx.WriteString("ok")
}