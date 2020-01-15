package controller

import (
	"github.com/gin-gonic/gin"
	"pea-web/api/plus"
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
	usr, err := service.UserService.Register(user.UserName, "", user.NickName, user.PassWord, user.RePassWord)
	if err != nil {
		plus.RespError(ctx, err)
		return
	}
	plus.RespSuccess(ctx, usr)
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
