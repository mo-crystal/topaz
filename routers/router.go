// @APIVersion 1.0
// @Title topaz user system
// @Contact github.com/mo-crystal
package routers

import (
	"topaz/controllers"

	"github.com/beego/beego"
	"github.com/beego/beego/context"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, CorsFilter)
	beego.Router("/", &controllers.ApiController{}, "GET:Status")
	beego.Router("/login", &controllers.ApiController{}, "POST:Login")
	beego.Router("/register", &controllers.ApiController{}, "POST:Register")

	// auth
	beego.Router("/pull-user", &controllers.ApiController{}, "POST:PullUser")
}

const (
	headerOrigin       = "Origin"
	headerAllowOrigin  = "Access-Control-Allow-Origin"
	headerAllowMethods = "Access-Control-Allow-Methods"
	headerAllowHeaders = "Access-Control-Allow-Headers"
)

func CorsFilter(ctx *context.Context) {
	origin := ctx.Input.Header(headerOrigin)
	ctx.Output.Header(headerAllowOrigin, origin)
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header(headerAllowMethods, "POST, GET, OPTIONS")
	ctx.Output.Header(headerAllowHeaders, "Content-Type, Authorization")
}
