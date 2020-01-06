package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-irain/logger"
)

type after struct {
	hooks []CustomHookFunc
}

func (mw *after) MiddleWare(egn *gin.Engine) {
	egn.Use(func(ctx *gin.Context) {
		func(ctx *gin.Context) {
			ctx.Next()
		}(ctx)

		// 执行自定义钩子方法
		for _, hook := range mw.hooks {
			err := hook(ctx)
			if err != nil {
				logger.Error(ctx.GetString("request_id"), err)
				ctx.Abort()
				break
			}
		}
		logResponse(ctx)  // 记录响应信息
		markResponse(ctx) // 标记响应信息
	})
}

// NewAfterMW 后置中间件
func NewAfterMW(hooks []CustomHookFunc) MiddleWarer {
	return &after{hooks: hooks}
}

// logResponse 响应信息记录到日志
func logResponse(ctx *gin.Context) {
	data, _ := ctx.Get("response_data")
	logger.Info(ctx.GetString("request_id"), "Reponse:", data)
}

// markResponse 标记响应
func markResponse(ctx *gin.Context) {
	st, _ := ctx.Get("http_stime")

	tf := "2006-01-02 15:04:05.000"
	sTime, _ := st.(time.Time)
	eTime := time.Now()
	duration := eTime.Sub(sTime)
	st2et := fmt.Sprintf("<%s ~ %s>", eTime.Format(tf), sTime.Format(tf))

	logger.Info(ctx.GetString("request_id"), "******* Duration:", duration.String(), st2et, "*******")
}
