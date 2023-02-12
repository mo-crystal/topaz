package controllers

import (
	"strconv"
	"topaz/database"
	"topaz/utils"
)

// POST /api/pull-user
//
// 拉取用户信息。
//
// 请求中需要指定回调地址。若地址在 target 列表中，并被标记为 Enabled = true，则会将用户信息推送至 CallbackUrl，并且向请求者返回一个标识符。
//
// 在请求者处，需接收来自本服务器的 POST 请求。将该 POST 请求中的标识符与本 API 返回的标识符对应，即可得到所拉取的用户信息。
func (c *ApiController) PullUser() {
	uid := c.Form("UserId")
	serverName := c.Form("Server")
	if serverName == "" {
		c.response(1, "server is nil")
		return
	}
	server := database.GetServerByName(serverName)
	password := c.Form("Password")
	signature := c.Form("Signature")
	if !utils.SignatureCheck(serverName+uid+password, signature, server.PublicKey) {
		c.response(2, "access deny")
		return
	}
	userId, err := strconv.Atoi(uid)
	if err != nil {
		c.response(3, "illegal user id")
		return
	}
	user := database.GetUser(userId)
	if password != user.Password {
		user.Mask()
		c.response(4, "wrong password", user)
		return
	}
	user.Password = ""
	c.response(0, "success", user)
}
