package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

/*
	接口成功
*/
func Success(ctx *gin.Context, data interface{}) {
	resp := make(map[string]interface{})
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
	resp["data"] = data
	ctx.JSON(http.StatusOK, resp)
}

/*
	接口失败
*/
func CheckError(ctx *gin.Context, errno string) {
	resp := make(map[string]interface{})
	resp["errno"] = errno
	resp["errmsg"] = RecodeText(errno)
	ctx.JSON(http.StatusOK, resp)
}

//异常捕获报错
func AnalysisError(ctx *gin.Context,err error,msg string)  {
	resp := make(map[string]interface{})
	resp["errno"] = "4890"
	resp["errmsg"] = errors.Wrap(err, msg).Error()
	ctx.JSON(http.StatusOK, resp)
}
