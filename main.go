package main

import (
    "github.com/astaxie/beego"
    "github.com/fanyang1988/goconfig"
    "github.com/fanyang1988/gologger"
)

var (
    ConfigManager = goconfig.New()
    LogManager    = gologger.New("logger", "./config/log.json", ConfigManager)
)

func init() {
}

func main() {
    defer ConfigManager.Close()
    defer LogManager.Close()

    beego.Router("/servers", &ServerInfoController{})
    beego.Run("0.0.0.0:7928")
}
