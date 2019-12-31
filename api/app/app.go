package app

import (
	"context"
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pea-web/api/controller"
	"pea-web/api/tools"
	"pea-web/cmd"
	"syscall"
	"time"
)

// 启动
func Start() {
	r := gin.Default()
	//日志中间件
	r.Use(ginzap.Ginzap(tools.NormalLogger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(tools.NormalLogger, true))
	//分组
	rg := r.Group("/api/v1.0")
	{
		rg.GET("/register", controller.Register)
	}

	host := fmt.Sprintf("%s:%s", cmd.Conf.HostURL, cmd.Conf.HostPort)
	srv := &http.Server{
		Addr:    host,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen:", err)
		}
	}()

	//优雅的关闭
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-c
	log.Println("shutDown server ....")

	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutDown ...", err)
	}
	log.Println("server Exiting")
}
