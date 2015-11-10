package routers

import (
	"github.com/astaxie/beego"
	"github.com/haoMonitor/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/project", &controllers.ProjectController{})
}
