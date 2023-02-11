package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"topaz/database"
)

// POST /api/pull-user
// 拉取用户信息。
// 请求中需要指定回调地址。若地址在 target 列表中，并被标记为 Enabled = true，则会将用户信息推送至 CallbackUrl，并且向请求者返回一个标识符。
// 在请求者处，需接收来自本服务器的 POST 请求。将该 POST 请求中的标识符与本 API 返回的标识符对应，即可得到所拉取的用户信息
func (c *ApiController) PullUser() {
	callbackUrl := c.Form("CallbackUrl")
	_uid := c.Form("UserId")
	if !CheckNotEmpty(callbackUrl, _uid) {
		c.response(1, "invalid form data")
		return
	}

	target := database.FindTarget(callbackUrl)
	if target == nil {
		c.response(2, "invalid callback url")
		return
	}

	uid, err := strconv.Atoi(_uid)
	if err != nil {
		c.response(3, "invalid user id")
		return
	}

	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	callbackData := Response{Code: 1, Msg: timestamp}
	requestedUser := database.GetUser(uid)
	if requestedUser != nil {
		callbackData.Code = 0
		callbackData.Data = requestedUser
	}

	callbackBytes, err := json.Marshal(callbackData)
	if err != nil {
		c.response(4, "internal error")
		return
	}

	_, err = http.Post(target.CallbackUrl, "application/json", bytes.NewReader(callbackBytes))
	if err != nil {
		c.response(5, "target server error")
		return
	}

	c.response(0, timestamp)
}
