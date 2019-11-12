package middleware

import "github.com/gin-gonic/gin"

// MiddleWarer 中间件接口
type MiddleWarer interface {
	MiddleWare(egn *gin.Engine)
}

// CustomHookFunc 自定义钩子方法
type CustomHookFunc func(ctx *gin.Context) error

// UseMiddleWare 执行中间件
func UseMiddleWare(egn *gin.Engine, bhooks []CustomHookFunc, ahooks []CustomHookFunc) {
	// 中间件组
	mw := []MiddleWarer{
		NewBeforeMW(bhooks),
		NewAfterMW(ahooks),
	}

	for _, m := range mw {
		m.MiddleWare(egn)
	}
}
