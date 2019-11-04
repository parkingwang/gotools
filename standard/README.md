# standard 工程项目标准规范包

## response 
响应消息规范化统一化

### 特性
* 有利于工程项目响应规范化、统一化
* 开箱即用
* 支持自定义链式调用
* 数据格式JSON、其他格式（推进中）
* 自定义返回结构（推进中）
* 自定义数据格式(推进中)

### 示例

```go
    // 开箱即用
    standard.RetSucc(ctx, []interface{}{})
    standard.RetFail(ctx, nil)

    // 链式调用
    standard.NewResponse(ctx).SetCode(1).SetMsg(msg).RetJSON()
```
