package controllers

import (
	"topaz/utils"

	"github.com/beego/beego"
)

type ApiController struct {
	beego.Controller
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// JSON 返回值的构造器，生成 1 个 JSON 响应并返回客户端
//
// 使用方法：
//
//	response(int) -> {"code": int, "message": "", "data": null}
//	response(int, string) -> {"code": int, "message": string, "data": null}
//	response(int, string, obj) -> {"code": int, "message": string, "data": obj}
func (c *ApiController) response(args ...interface{}) {
	resp := Response{Code: 0}
	switch len(args) {
	case 3:
		resp.Data = args[2]
		fallthrough
	case 2:
		resp.Msg = args[1].(string)
		fallthrough
	case 1:
		resp.Code = args[0].(int)
	}

	if resp.Code >= 100 {
		c.Ctx.ResponseWriter.WriteHeader(resp.Code)
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *ApiController) Form(key string) string {
	return c.Ctx.Request.Form.Get(key)
}

// 校验任意个 string，其中只要包含空字符串，就会返回 false
func CheckNotEmpty(str ...string) bool {
	for _, s := range str {
		if s == "" {
			return false
		}
	}
	return true
}

// @Title 获取服务器状态
// @Description 获取服务器状态，确认版本号
// @Success 200 {object} controllers.Response
// @router /api/ [get]
func (c *ApiController) Status() {
	c.response(0, "topaz server: "+utils.VERSION)
}
