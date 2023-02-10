package main

import (
	"topaz/database"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.Init()

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
