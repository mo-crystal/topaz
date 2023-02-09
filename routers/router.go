package routers

import (
	"topaz/controllers"

	"github.com/beego/beego"
)

func init() {
	beego.Router("/api/", &controllers.ApiController{}, "GET:Status")
	beego.Router("/api/login", &controllers.ApiController{}, "POST:Login")
	beego.Router("/api/register", &controllers.ApiController{}, "POST:Register")
}
