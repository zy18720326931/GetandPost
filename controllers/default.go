package controllers
//config配置
import (
	"Proinit0/Porseron"
	"Proinit0/db_mysql"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
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
	var poser Porseron.Poser1
	strbyte , err:= ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		fmt.Println("解析错误")
		return
	}
	err2:=json.Unmarshal(strbyte,&poser)
	if err2 !=nil{
		fmt.Println("传输错误")
		return
	}

	err3,id:=db_mysql.Set(poser)
	if err3 != nil {
		fmt.Println(err.Error())
		fmt.Println("未插入")
		return
	}
    fmt.Println(id)
	fmt.Println(poser.Name)
	fmt.Println(poser.Sex)
	fmt.Println(poser.Age)
	c.Ctx.WriteString("成功")
}
//postman :用于模拟http各种类型请求的一个工具，浏览器