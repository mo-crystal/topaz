package controllers

import (
	"strconv"
	"topaz/database"
	"topaz/utils"
)

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
	if user == nil {
		c.response(4, "user not exists")
		return
	}

	if password != user.Password {
		user.Mask()
		c.response(0, "wrong password", user)
		return
	}
	user.Password = ""
	c.response(0, "success", user)
}
