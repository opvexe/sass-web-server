package handler

import (
	"context"
	user "pea-web/service/user-service/proto/user"
)

type User struct{}

//创建用户
func (e *User) Create(ctx context.Context, req *user.Request, rsp *user.Response) error {

	return nil
}

//验证用户
func (e *User) Auth(ctx context.Context, req *user.Request, rsp *user.Token) error {

	return nil
}
