# gin框架使用过程中抽取的复用包

## middleware 


### 特性
* 封装前置和后置中间钩子:标记请求、记录请求、记录响应信息
* 支持添加自定义前置或后置处理方法，要求是`[]CustomHookFunc`函数类型
* 持续优化推进中...


### 示例
```go
package main 

import (
    ... 
    "github.com/parkingwang/gotools/gin/middleware"
    "github.com/parkingwang/gotools/standard"
)

var beforehookFs = []middleware.CustomHookFunc{
    // 只允许JSON格式数据
    func(ctx *gin.Context) {
        ctp := ctx.ContentType()
        if ctp != "" && ctp != gin.MIMEJSON {
            msg := "仅支持JSON数据格式"
            standard.NewResponse(ctx).SetCode(1).SetMsg(msg).RetJSON()
            ctx.Abort()
        }
    },
}
func main(){
    engine := gin.Default()
    middleware.UseMiddleWare(engine, beforehookFs, nil) // 中间件

    ...
}

```
