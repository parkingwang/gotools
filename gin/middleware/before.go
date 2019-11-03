package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-irain/logger"
	"github.com/parkingwang/gotools/funcs"
	"github.com/parkingwang/gotools/standard"
)

type before struct{}

func (mw *before) MiddleWare(egn *gin.Engine) {
	egn.Use(func(ctx *gin.Context) {
		markRequest(ctx) // 标记请求
		logRequest(ctx)  // 请求信息记录
		chkJSON(ctx)     // json的http请求体检查

		func(ctx *gin.Context) {
			ctx.Next()
		}(ctx)
	})
}

// NewBeforeMW 前置中间件
func NewBeforeMW() MiddleWarer {
	return &before{}
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

	RequestID, _ := ctx.Get("request_id")
	ctx.Request.ParseMultipartForm(1024)
	logger.Info(RequestID, "\n\n************ Client : ", ctx.Request.RemoteAddr, " [", ctx.Request.Method, "] ", ctx.Request.URL.Path, "************")
	logger.Info(RequestID, "URL: ", ctx.Request.URL, "ContentType: ", ctx.ContentType())
	logger.Info(RequestID, " PostForm:", ctx.Request.PostForm, "JSON:", string(bodyBytes))
}

// chkJSON 检查json形式的http请求
func chkJSON(ctx *gin.Context) {
	ctp := ctx.ContentType()
	if ctp != "" && ctp != gin.MIMEJSON {
		msg := "仅支持JSON数据格式"
		standard.NewResponse(ctx).SetCode(1).SetMsg(msg).RetJSON()

		ctx.Abort()
	}
}
