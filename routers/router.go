package routers

import (
	"topaz/controllers"

	"github.com/beego/beego"
	"github.com/beego/beego/context"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, CorsFilter)
	beego.Router("/api/", &controllers.ApiController{}, "GET:Status")
	beego.Router("/api/login", &controllers.ApiController{}, "POST:Login")
	beego.Router("/api/register", &controllers.ApiController{}, "POST:Register")
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
