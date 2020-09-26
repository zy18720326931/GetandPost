package main

import (
	_ "Proinit0/db_mysql"
	_ "Proinit0/routers"

	"github.com/astaxie/beego"
)

func main(){
      beego.Run()
}

