package controller

import (
	"github.com/gin-gonic/gin"
	"pea-web/api/service"
	"pea-web/api/tools"
)

/*
	注册
*/
func Register(ctx *gin.Context) {
	var user struct {
		UserName   string `json:"username"`
		PassWord   string `json:"password"`
		RePassWord string `json:"rePassword"`
		NickName   string `json:"nickname"`
		Ref        string `json:"ref"`
	}
	err := ctx.Bind(&user)
	if err != nil {
		return
	}
	usr, err := service.UserService.Register(user.UserName, "", user.NickName, user.PassWord, user.RePassWord)
	if usr == nil {
		tools.CheckError(ctx, err, "参数错误")
	}
	tools.Success(ctx, usr)
}