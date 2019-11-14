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

### 响应格式

字段 | 类型 | 必须 | 备注
---|---|---|---
status | uint | 是 | 响应码: 1-成功;2-失败;其他-自定义
msg | string | 是 | 响应消息
data | object | 否 | 响应数据对象


### API


### 示例

```go
    // 开箱即用
    standard.RetSucc(ctx, []interface{}{})
    standard.RetFail(ctx, nil)

    // 链式调用
    standard.NewResponse(ctx).Status(1).Msg(msg).Data([]interface{}{}).RetJSON()
```


```json
{
    "status":1,
    "msg":"操作成功",
    "data": {
        ...
    }
}
```