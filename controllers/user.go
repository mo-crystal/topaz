package controllers

import (
	"strconv"
	"topaz/database"
)

func (c *ApiController) GetUser() *database.User {
	data := c.GetSession("user")
	if data == nil {
		return nil
	}

	uid, ok := data.(int)
	if !ok {
		panic("session data error")
	}

	u := database.GetUser(uid)
	u.Password = ""
	return u
}

func (c *ApiController) SetUser(uid int) {
	c.SetSession("user", uid)
}

func (c *ApiController) RemoveUser() {
	c.DelSession("user")
}

func (c *ApiController) Register() {
	name := c.Form("name")
	password := c.Form("password")
	if !CheckNotEmpty(name, password) {
		c.response(1, "Data error.")
		return
	}

	user := &database.User{
		Name:     name,
		Password: password,
		Admin:    false,
		Banned:   false,
	}
	database.SetUser(user)
	c.response(0, "", user)
}

func (c *ApiController) Login() {
	idstr := c.Form("id")
	password := c.Form("password")
	if !CheckNotEmpty(idstr, password) {
		c.response(1, "Missing parameters.")
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.response(2, "Invalid uid.")
		return
	}

	user := database.GetUser(id)
	if user == nil {
		c.response(3, "User not exists.")
		return
	}

	if user.Password != password {
		c.response(4, "Wrong password.")
		return
	}

	c.SetUser(user.Id)
	user.Password = ""
	c.response(0, "Login accomplished.", user)
}

func (c *ApiController) Account() {
	user := c.GetUser()
	if user == nil {
		c.response(1, "Please login.")
		return
	}

	user.Password = ""
	c.response(0, "", user)
}

func (c *ApiController) SignOut() {
	c.RemoveUser()
	c.response()
}
