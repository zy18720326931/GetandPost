package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func init()  {
	config:=beego.AppConfig
	dbDriver :=config.String("db_driverName")
	dbUser :=config.String("db_user")
	dbpassword:=config.String("db_password")
	db_ip:=config.String("db_ip")
	db_name:=config.String("db_name")
	connurls:=dbUser+":"+dbpassword+"@tcp("+db_ip+")/"+db_name+"?charset=utf8"
	db,err:=sql.Open(dbDriver,connurls)
	if err!=nil {
		fmt.Println("连接出错")
		return
	}
	DB=db
	fmt.Println("连接成功")
}