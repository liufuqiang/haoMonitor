package models

import (
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id     int64
	Name   string
	Status int8
}

func RegisterDB() {
	orm.RegisterModel(new(Project))
	//	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/haoMonitor?charset=utf8") //密码为空格式
}
