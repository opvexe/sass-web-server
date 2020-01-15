package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pea-web/api/plus"
	"pea-web/api/tools"
)

//崩溃拦截
func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if tools.NormalLogger != nil {
					tools.NormalLogger.Error("异常崩溃" + fmt.Sprint(err))
				}
				plus.RespError(context, plus.PE_ServerError)
				return
			}
		}()
		context.Next()
	}
}
