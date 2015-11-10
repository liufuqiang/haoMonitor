package main

import (
	"github.com/astaxie/beego"
	_ "github.com/haoMonitor/routers"
	"github.com/haomonitor/models"
)

func main() {
	models.RegisterDB()
	beego.Run()
}
