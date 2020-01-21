package handler

import (
	"context"
	user "pea-web/service/user-service/proto/user"
	"pea-web/service/user-service/service"
)

type User struct{}

//注册用户
func (e *User) MicroRegist(ctx context.Context, req *user.RegistRequest, rsp *user.RegistResponse) error {

	_, err := service.UserService.Register(req.UserName, req.Email, req.NickName, req.Password, req.PasswordAgain)
	if err != nil {
		rsp.Code = 450
		rsp.Message = "注册用户失败"
		return err
	}
	rsp.Code = 200
	rsp.Message = "操作成功"
	return nil
}

//验证用户
func (e *User) MicroLogin(ctx context.Context, req *user.LoginRequest, rsp *user.LoginResponse) error {

	usr, err := service.UserService.Login(req.UserName, req.Password)
	if err != nil {
		rsp.Code = 451
		rsp.Message = "用户登录失败"
	}
	rsp.Code = 200
	rsp.Message = "操作成功"
	var u user.Account
	u.Id = int32(usr.ID)
	u.Nickname = usr.Nickname
	u.Email = usr.Email
	u.Avatar = usr.Avatar
	u.Status = int32(usr.Status)
	u.Roles = usr.Roles
	u.Type = int32(usr.Type)
	rsp.Data = &u
	return nil
}
