package main

import (
	"os"
	"topaz/database"
	_ "topaz/routers"

	"github.com/beego/beego"
)

func main() {
	os.Mkdir("data", os.ModePerm)
	database.Init()

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
