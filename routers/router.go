package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"topaz/controllers"
)

func init() {
	beego.Router("/api/", &controllers.ApiController{}, "GET:Status")
	beego.Router("/api/login", &controllers.ApiController{}, "POST:Login")
	beego.Router("/api/register", &controllers.ApiController{}, "POST:Register")
}
