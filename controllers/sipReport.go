package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type SipController struct {
	web.Controller
}

func (this *SipController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "sip/index.tpl"
	this.Render()
}
