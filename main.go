package main

import (
    "github.com/astaxie/beego"
    "github.com/fanyang1988/goconfig"
    "github.com/fanyang1988/gologger"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Ctx.WriteString("hello world")
}

var (
    ConfigManager = goconfig.NewConfig()
    LogManager    = log.NewLog("logger", "./config/log.json", ConfigManager)
)

func init() {
}

func main() {
    defer ConfigManager.Close()
    defer LogManager.Close()

    server_info := &ServerInfoController{}

    beego.Router("/", &MainController{})
    beego.Router("/servers", server_info)
    beego.Run("0.0.0.0:7928")
}
