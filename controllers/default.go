package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
