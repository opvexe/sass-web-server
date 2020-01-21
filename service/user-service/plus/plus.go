package plus

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//响应数据
type Response struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"omitempty"`
	ERR        error  `json:"omitempty"`
}

// 自定义错误
func (r *Response) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

//解包响应错误
func UnWrapResponse(err error) *Response {
	if v, ok := err.(*Response); ok {
		return v
	}
	return nil
}

//响应成功  200
func RespSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, "操作成功", 200, v)
}

//响应错误
func RespError(c *gin.Context, err error, status ...int) {
	var resp *Response
	if err != nil {
		if e, ok := err.(*Response); ok {
			resp = e
		} else {
			resp = UnWrapResponse(Wrap500Response(err))
		}
	} else {
		resp = UnWrapResponse(PE_ServerError)
	}
	if len(status) > 0 {
		resp.StatusCode = status[0]
	}
	if err := resp.ERR; err != nil {
		c.AbortWithStatusJSON(resp.StatusCode, gin.H{
			"code":    resp.Code,
			"message": resp.ERR,
		})
		return
	}
	c.AbortWithStatusJSON(resp.StatusCode, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
	})
}

//[解析错误] 也属于400错误
func Parse(c *gin.Context, obj interface{}) error {
	if err := c.Bind(obj); err != nil {
		return Wrap400Response(err, "解析请求参数发生错误")
	}
	return nil
}

//状态码为400请求错误
func Wrap400Response(err error, msg ...string) error {
	m := "请求发生错误"
	if len(msg) > 0 {
		m = msg[0]
	}
	return WrapResponse(err, 400, m, 400)
}

//服务器500响应错误
func Wrap500Response(err error, msg ...string) error {
	m := "服务器发生错误"
	if len(msg) > 0 {
		m = msg[0]
	}
	return WrapResponse(err, 500, m, 500)
}

//公共方法
func WrapResponse(err error, code int, msg string, status ...int) error {
	resp := &Response{
		Code:    code,
		Message: msg,
		ERR:     err,
	}
	if len(status) > 0 {
		resp.StatusCode = status[0]
	}
	return resp
}

//新公共方法
func NewWrapResponse(code int, msg string, status ...int) error {
	resp := &Response{
		Code:    code,
		Message: msg,
	}
	if len(status) > 0 {
		resp.StatusCode = status[0]
	}
	return resp
}

//状态码为400请求错误
func NewWrap400Response(msg string) error {
	return NewWrapResponse(400, msg, 400)
}

//服务器500响应错误
func NewWrap500Response(msg string) error {
	return NewWrapResponse(500, msg, 500)
}

//响应JSON
func ResJSON(c *gin.Context, status int, msg string, code int, v interface{}) {
	c.AbortWithStatusJSON(status, gin.H{
		"code":    code,
		"message": msg,
		"data":    v,
	})
}
