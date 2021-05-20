package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type UserInfoController struct {
	web.Controller
}

func (this *UserInfoController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "user/index.tpl"
	this.Render()
}
