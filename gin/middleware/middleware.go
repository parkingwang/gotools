package middleware

import "github.com/gin-gonic/gin"

// MiddleWarer 中间件接口
type MiddleWarer interface {
	MiddleWare(egn *gin.Engine)
}

// UseMiddleWare 执行中间件
func UseMiddleWare(egn *gin.Engine) {
	// 中间件组
	mw := []MiddleWarer{
		NewBeforeMW(),
		NewAfterMW(),
	}

	for _, m := range mw {
		m.MiddleWare(egn)
	}
}
