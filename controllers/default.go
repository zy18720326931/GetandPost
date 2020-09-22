package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "zy18720326931@126.com"
	c.TplName = "index.tpl"
}
