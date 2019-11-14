package standard

import (
	"net/http"
)

const (
	// FAIL 失败状态码
	FAIL uint = iota
	// SUCC 成功状态码
	SUCC
)

// Msgs 提示信息
var Msgs = map[uint]string{
	FAIL: "操作失败",
	SUCC: "操作成功",
}

// JSONer json响应接口
type JSONer interface {
	JSON(code int, obj interface{})
	Set(key string, value interface{})
}

// Response 响应消息结构体
type Response struct {
	jsoner JSONer
	msgs   map[uint]string
	status uint
	msg    string
	data   interface{}
}

// Status 设置响应码
func (rsp *Response) Status(status uint) *Response {
	rsp.status = status
	return rsp
}

// Msg 设置响应消息
func (rsp *Response) Msg(msg string) *Response {
	rsp.msg = msg
	return rsp
}

// Data 设置响应数据
func (rsp *Response) Data(data interface{}) *Response {
	rsp.data = data
	return rsp
}

// Msgs 提示信息扩展
func (rsp *Response) Msgs(msgs map[uint]string) *Response {
	rsp.msgs = msgs
	return rsp
}

// Raw 链式设置
func (rsp *Response) Raw(mix interface{}) *Response {
	switch i := mix.(type) {
	case uint:
		rsp.Status(i)
	case error:
		rsp.Msg(i.Error())
	case string:
		rsp.Msg(i)
	case map[uint]string:
		rsp.Msgs(i)
	default:
		rsp.Data(i)
	}

	return rsp
}

// RetJSON 返回json消息
func (rsp *Response) RetJSON() {
	data := map[string]interface{}{
		"status": rsp.status,
		"msg":    rsp.msg,
		"data":   rsp.data,
	}
	rsp.jsoner.Set("response_data", data)
	rsp.jsoner.JSON(http.StatusOK, data)
}

// NewResponse 生成一个响应结构体指针
func NewResponse(j JSONer) *Response {
	return &Response{
		jsoner: j,
		msgs:   Msgs,
	}
}

// RetSucc 成功响应
func RetSucc(j JSONer, data interface{}) {
	NewResponse(j).Status(SUCC).Msg(Msgs[SUCC]).Data(data).RetJSON()
}

// RetFail 失败响应
func RetFail(j JSONer, data interface{}) {
	NewResponse(j).Status(FAIL).Msg(Msgs[FAIL]).Data(data).RetJSON()
}

// RetMix 多个参数混合构建响应
func RetMix(j JSONer, args ...interface{}) {
	rsp := NewResponse(j)
	for _, d := range args {
		rsp.Raw(d)
	}
	rsp.RetJSON()
}

// RetMixSucc 多个参数混合构建成功响应
func RetMixSucc(j JSONer, args ...interface{}) {
	rsp := NewResponse(j).Msg(Msgs[SUCC])
	for _, d := range args {
		rsp.Raw(d)
	}
	rsp.Status(SUCC).RetJSON()
}

// RetMixFail 多个参数混合构建失败响应
func RetMixFail(j JSONer, args ...interface{}) {
	rsp := NewResponse(j).Msg(Msgs[FAIL])
	for _, d := range args {
		rsp.Raw(d)
	}
	rsp.Status(FAIL).RetJSON()
}
