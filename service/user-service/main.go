package main

import (
	"fmt"
	"github.com/micro/go-micro/registry"
	etcd "github.com/micro/go-plugins/registry/etcdv3"
	"log"
	"net/http"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"pea-web/service/user-service/db"
	"pea-web/service/user-service/model"
	"pea-web/service/user-service/handler"
	user "pea-web/service/user-service/proto/user"
	"time"
)

func main() {
	//创建数据库连接，程序退出时断开连接
	db, err := db.CreateConnection()
	defer db.Close()

	if err!=nil{
		log.Fatalf("数据库连接失败:%v",err)
	}
	//每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(new(model.User))
	db.AutoMigrate(new(model.UserToken))

	//创建微服务流程
	regist := etcd.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379", "http://127.0.0.1:2380",
		}
		op.Timeout = 5 * time.Second	//5秒超时
	})
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Registry(regist),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),// 基于 prometheus 采集监控指标数据
	)
	srv.Init()

	// 注册处理器
	if err :=user.RegisterUserHandler(srv.Server(), new(handler.User));err!=nil{
		fmt.Println(err)
	}

	// 采集监控数据
	prometheusBoot()

	// 启动用户服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

//启动 HTTP 服务监听 Prometheus 客户端监控数据采集
func prometheusBoot()  {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}
