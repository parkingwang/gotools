package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-irain/logger"
)

type after struct{}

func (mw *after) MiddleWare(egn *gin.Engine) {
	egn.Use(func(ctx *gin.Context) {
		func(ctx *gin.Context) {
			ctx.Next()
		}(ctx)

		logResponse(ctx)  // 记录响应信息
		markResponse(ctx) // 标记响应信息
	})
}

// NewAfterMW 后置中间件
func NewAfterMW() MiddleWarer {
	return &after{}
}

// logResponse 响应信息记录到日志
func logResponse(ctx *gin.Context) {
	logid, _ := ctx.Get("logid")
	data, _ := ctx.Get("response_data")
	logger.Info(logid, "Reponse:", data)
}

// markResponse 标记响应
func markResponse(ctx *gin.Context) {
	logid, _ := ctx.Get("logid")
	sTime, _ := ctx.Get("http_stime")
	eTime := time.Now()
	duration := eTime.Sub(sTime.(time.Time))
	logger.Info(logid, "http[stime:", sTime, "~ etime:", eTime, "]\n************ Cost Duration : ", duration.String(), " ************\r\n")
}
