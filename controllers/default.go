package controllers
//config配置
import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	UserName:=c.Ctx.Input.Query("user")
	UserAge:=c.Ctx.Input.Query("age")
	if  UserAge !="20"||UserName !="zy"{
		c.Ctx.ResponseWriter.Write([]byte("输入不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("Wlecom"))
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "zy18720326931@126.com"
	c.TplName = "index.tpl"
}


func (c *MainController)Post(){
	UserName:=c.Ctx.Request.FormValue("user")
	UserAge:=c.Ctx.Request.FormValue("age")
    Uaersex:=c.Ctx.Request.FormValue("sex")
	if UserName!="zy"||UserAge !="20"||Uaersex !="main" {
		c.Ctx.WriteString("hellow!!")
		return
	}
	c.Ctx.WriteString("chenggong")
}
//postman :用于模拟http各种类型请求的一个工具，浏览器