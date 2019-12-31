package main

import (
	"fmt"
	"pea-web/api/app"
	"pea-web/api/tools"
	"pea-web/cmd"
)

// go run main.go -c=./pea.yaml
func main() {
	//初始化配置
	if err := cmd.DebugStart(); err != nil {
		fmt.Println("start err :", err)
		panic(err)
	}
	//启动日志
	tools.InitWithLog(cmd.Conf.DevEnv, cmd.Conf.LoggerFile)
	//启动web
	app.Start()
}
