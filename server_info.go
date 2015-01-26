package main

import (
    "github.com/astaxie/beego"
)

var (
    logger = LogManager.GetLogger("serverInfo")
)

func init() {
    ConfigManager.Reg("server_info", "./config/server_info.json", true)
}

type ServerInfoController struct {
    beego.Controller
}

func (self *ServerInfoController) Prepare() {
    name := self.GetString("name")
    if name == "" {
        logger.Error("Get Info Param Err")
        self.StopRun()
    }
}

func (self *ServerInfoController) Get() {
    name := self.GetString("name")
    info := ConfigManager.Get("server_info")

    server_info := info.Get(name)
    str, err := server_info.Encode()
    if err == nil {
        logger.Info("GetInfo %s", str)
        self.Ctx.WriteString(string(str))
    } else {
        logger.Error("No Info %s", name)
        self.Ctx.WriteString("Error")
    }

}

func (self *ServerInfoController) Finish() {
    logger.Debug("Finish")
}
