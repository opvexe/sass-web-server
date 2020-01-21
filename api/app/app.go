package app

import (
	"context"
	"fmt"
	"github.com/Allenxuxu/microservices/lib/tracer"
	"github.com/Allenxuxu/microservices/lib/wrapper/tracer/opentracing/gin2micro"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pea-web/api/controller"
	"pea-web/api/middleware"
	"pea-web/api/plus"
	"pea-web/cmd"
	"syscall"
	"time"
)

// 启动
func Start() {
	//jaeger分布式链路追踪
	gin2micro.SetSamplingFrequency(50) //设置采样率
	t, io, err := tracer.NewTracer("pea.web.gin.api", "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t) //设置全局追踪

	host := fmt.Sprintf("%s:%s", cmd.Conf.HostURL, cmd.Conf.HostPort)
	srv := &http.Server{
		Addr:    host,
		Handler: initWithApp(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("服务监听:", err)
		}
	}()

	//优雅的关闭
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-c
	fmt.Println("服务正在关闭中...")

	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("服务关闭:", err)
	}
	fmt.Println("服务退出...")
}

func initWithApp() *gin.Engine {
	app := gin.Default()

	////日志中间件
	//r.Use(ginzap.Ginzap(tools.NormalLogger, time.RFC3339, true))
	//r.Use(ginzap.RecoveryWithZap(tools.NormalLogger, true))

	//错误拦截
	app.Use(middleware.RecoveryMiddleware())

	//jaeger链路追踪
	app.Use(gin2micro.TracerWrapper)

	initUserRouter(app)

	app.NoRoute(func(context *gin.Context) {
		plus.RespError(context, plus.PE_NotFoundRouter)
	})

	return app
}

func initUserRouter(app *gin.Engine) {

	rg := app.Group("/api")

	{
		rg.GET("/register", controller.Register)
	}
}
