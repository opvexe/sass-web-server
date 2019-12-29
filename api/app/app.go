package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/gin-contrib/zap"
	"os"
	"os/signal"
	"pea-web/api/controller"
	"pea-web/cmd"
	"pea-web/tools"
	"syscall"
	"time"
)

// 启动
func Start() {
	r := gin.Default()
	//日志中间件
	r.Use(ginzap.Ginzap(tools.NormalLogger,time.RFC3339,true))
	r.Use(ginzap.RecoveryWithZap(tools.NormalLogger,true))
	//分组
	rg := r.Group("/api/v1.0")
	{
		rg.GET("/register", controller.Register)
	}

	host := fmt.Sprintf("%s:%s", cmd.Conf.HostURL, cmd.Conf.HostPort)
	if err := r.Run(host); err != nil {
		panic(err)
	}

	//优雅的关闭
	c := make(chan os.Signal)
	//指定关闭的信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		s := <-c
		logrus.Infof("got signal [%s], exiting now", s)
		//关闭数据库
		if err := cmd.DB.Close(); err != nil {
			logrus.Error("mysql service closed failed:", err)
		}
		//关闭reids
		if err := cmd.RDS.Close(); err != nil {
			logrus.Error("redis service closed failed:", err)
		}
		logrus.Infof("程序退出")
		os.Exit(0)
	}()
}
