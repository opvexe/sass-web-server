package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pea-web/tools"
	"time"
)

type DataAnalysisModel struct {
	RequstTime string `json:"datetime"`
	RequestURL  string `json:"url"`
}

var DataAnalysisCh chan DataAnalysisModel

func init()  {
	//初始化管道
	DataAnalysisCh = make(chan DataAnalysisModel,1000)
	go HanleChannel()
}

// 中间件
func CustomMiddleWare() func(*gin.Context) {
	return func(context *gin.Context) {
		//拦截崩溃
		defer func() {
			if err:=recover();err!=nil{
				if tools.NormalLogger!=nil{
					tools.NormalLogger.Error("API异常"+fmt.Sprint(err))
				}
				tools.AnalysisError(context, errors.New(fmt.Sprint(err)), "异常消息")
				return
			}
		}()

		//上报用户活跃度
		go SendDataAnalysis(context)
	}
}


//上报用户活跃度
func SendDataAnalysis(ctx *gin.Context)  {
	//异常捕获
	defer func() {
		if err:=recover();err!=nil{
			tools.AnalysisError(ctx, errors.New(fmt.Sprint(err)), "异常消息")
			return
		}
	}()
	var d DataAnalysisModel
	d.RequstTime = time.Now().Format("2006-01-02 15:04:05")
	d.RequestURL = string(ctx.Request.URL.Path)
	//获取用户id
	DataAnalysisCh <- d
}

func HanleChannel()  {
	for{
		select {
		case c := <- DataAnalysisCh:
		fmt.Println("上报数据",c)
		
		default:
			time.Sleep(10*time.Second)
		}
	}
}