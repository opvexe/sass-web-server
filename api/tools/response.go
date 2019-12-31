package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

/*
 请求成功
*/
func Success(ctx *gin.Context, data interface{}) {
	resp := make(map[string]interface{})
	resp["errno"] = 0
	resp["errmsg"] = "SUCCESS"
	resp["data"] = data
	ctx.JSON(http.StatusOK, resp)
}

/*
 请求失败
*/
func CheckError(ctx *gin.Context, err error, msg string) {
	resp := make(map[string]interface{})
	resp["errno"] = 1
	resp["errmsg"] = errors.Wrap(err, msg).Error()
	ctx.JSON(http.StatusOK, resp)
}
