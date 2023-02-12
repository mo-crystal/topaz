package main

import (
	"fmt"
	"os"
	"topaz/database"
	_ "topaz/routers"
	"topaz/utils"

	"github.com/beego/beego"
)

func main() {
	os.Mkdir("data", os.ModePerm)
	utils.OverrideFile(fmt.Sprintf("%d", os.Getpid()), "data/pid")

	database.Init()

	beego.BConfig.WebConfig.Session.SessionOn = true
	fmt.Println("1")
	beego.Run()
}
