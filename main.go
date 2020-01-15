package main

import (
	"pea-web/api/app"
	"pea-web/api/tools"
	"pea-web/cmd"
)

func main() {
	//初始化配置
	if err := cmd.DebugStart(); err != nil {
		panic(err)
	}
	//启动日志
	tools.InitWithLog(cmd.Conf.DevEnv, cmd.Conf.LoggerFile)
	app.Start()
}
