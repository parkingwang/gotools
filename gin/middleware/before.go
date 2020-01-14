package middleware

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-irain/logger"
	"github.com/parkingwang/gotools/funcs"
)

type before struct {
	hooks []CustomHookFunc
}

func (mw *before) MiddleWare(egn *gin.Engine) {
	egn.Use(func(ctx *gin.Context) {
		markRequest(ctx) // 标记请求
		logRequest(ctx)  // 请求信息记录
		// 执行自定义钩子方法
		for _, hook := range mw.hooks {
			err := hook(ctx)
			if err != nil {
				logger.Error(ctx.GetString("request_id"), err)
				ctx.Abort()
				break
			}
		}

		func(ctx *gin.Context) {
			ctx.Next()
		}(ctx)
	})
}

// NewBeforeMW 前置中间件
func NewBeforeMW(hooks []CustomHookFunc) MiddleWarer {
	return &before{hooks: hooks}
}

// markRequest 标记请求
func markRequest(ctx *gin.Context) {
	ctx.Set("http_stime", time.Now())
	ctx.Set("request_id", funcs.RequestID())
}

// logRequest 请求信息记录到日志
func logRequest(ctx *gin.Context) {
	var bodyBytes []byte
	if ctx.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	RequestID := ctx.GetString("request_id")
	ctx.Request.ParseMultipartForm(1024)
	logger.Info(RequestID, "******* Client : ", ctx.Request.RemoteAddr, " [", ctx.Request.Method, "] ", ctx.Request.URL.Path, "*******")
	logger.Info(RequestID, "URL: ", ctx.Request.URL, "ContentType: ", ctx.ContentType())
	if !strings.Contains(ctx.GetHeader("Content-Disposition"), "file") {
		logger.Info(RequestID, "PostForm:", ctx.Request.PostForm, "JSON:", string(bodyBytes))
	}
}
