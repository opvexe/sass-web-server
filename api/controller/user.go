package controller

import (
	"errors"
	"github.com/Allenxuxu/microservices/lib/wrapper/tracer/opentracing/gin2micro"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"pea-web/api/plus"
	upb "pea-web/api/proto/user"
	"pea-web/api/service"
)

/*
	注册
*/
func Register(ctx *gin.Context) {
	var user struct {
		UserName   string `json:"username" form:"username"`
		PassWord   string `json:"password" form:"password"`
		RePassWord string `json:"rePassword" form:"rePassword"`
		NickName   string `json:"nickname" form:"nickname"`
	}
	err := ctx.Bind(&user)
	if err != nil {
		return
	}
	c, ok := gin2micro.ContextWithSpan(ctx)
	if !ok {
		plus.RespError(ctx, errors.New("Grpc调用失败"))
		return
	}

	regist := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379", "http://127.0.0.1:2380",
		}
	})
	srv := micro.NewService(
		micro.Registry(regist),
	)
	client := upb.NewUserService("go.micro.srv.user", srv.Client())
	if resp, err := client.MicroRegist(c, &upb.RegistRequest{
		UserName:      user.UserName,
		Email:         "",
		NickName:      user.NickName,
		Password:      user.PassWord,
		PasswordAgain: user.RePassWord,
	}); err == nil {
		plus.RespSuccess(ctx, resp)
		return
	}
	plus.RespError(ctx, errors.New("Grpc调用失败"))
}

//用户密码登录
func Login(ctx *gin.Context) {
	var user struct {
		UserName string `json:"username" form:"username"`
		PassWord string `json:"password" form:"password"`
	}
	err := ctx.Bind(&user)
	if err != nil {
		return
	}
	usr, err := service.UserService.Login(user.UserName, user.PassWord)
	if err != nil {
		plus.RespError(ctx, err)
		return
	}
	plus.RespSuccess(ctx, usr)
}
