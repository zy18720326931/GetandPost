package db_mysql

import (
	"Proinit0/Porseron"
	"fmt"
)

func Set(	pose  Porseron.Poser1)(error,int64)  {
	 result,err:=DB.Exec("insert  into  Userdata(id,username,brithdy,adress,age,sex)values (?,?,?,?,?,?)",1,pose.Name,pose.Borthdy,pose.Adress,pose.Age,pose.Sex)
	if  err !=nil {
		fmt.Println(err.Error())
		fmt.Println("插入数据出错")
		return err, 0
	}
    id ,err2:=result.RowsAffected()
	if  err2 !=nil {
		fmt.Println(err2.Error())
		fmt.Println("读取插入数据出错")
		return err2,0
	}
	return nil, id
}