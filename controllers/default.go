package controllers

import (
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

func CheckNotEmpty(str ...string) bool {
	for _, s := range str {
		if s == "" {
			return false
		}
	}
	return true
}

func (c *ApiController) Status() {
	c.response(0, "v1.0")
}
