package main

import (
	"fmt"
	"os"
	"topaz/database"
	_ "topaz/routers"

	"github.com/beego/beego"
)

func main() {
	os.Mkdir("data", os.ModePerm)
	database.Init()

	beego.BConfig.WebConfig.Session.SessionOn = true
	fmt.Println("1")
	beego.Run()
}
