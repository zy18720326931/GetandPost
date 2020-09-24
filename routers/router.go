package routers

import (
	"Proinit0/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/lergin", &controllers.MainController{})
}
//路由
