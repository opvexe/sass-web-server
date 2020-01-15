package controller

import (
	"github.com/gin-gonic/gin"
	"pea-web/api/model"
	"pea-web/api/service"
	"pea-web/api/tools"
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
	if usr == nil {
		tools.CheckError(ctx, err, "参数错误")
		return
	}
	tools.Success(ctx, usr)
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

}

//登录成功后生成token
func GenerateToken(user *model.User, ref string) {

}
