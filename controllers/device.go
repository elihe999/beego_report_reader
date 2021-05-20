package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type DeviceController struct {
	web.Controller
}

func (this *DeviceController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "device/index.tpl"
	this.Render()
}
