package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/haomonitor/models"
)

type ProjectController struct {
	beego.Controller
}

func (c *ProjectController) Get() {

	o := orm.NewOrm()

	project := &models.Project{}
	project.Name = "张三"
	project.Status = 1
	fmt.Println(o.Insert(project))

	c.Data["Website"] = "haosou.com"
	c.Data["Email"] = "liufuqiang@360.cn"
	c.TplNames = "projectAdd.tpl"
}
