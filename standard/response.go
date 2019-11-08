package standard

import (
	"net/http"
)

// JSONer json响应接口
type JSONer interface {
	JSON(code int, obj interface{})
	Set(key string, value interface{})
}

// Response 响应消息结构体
type Response struct {
	jsoner  JSONer
	code    int
	message string
	data    interface{}
}

// Code 设置响应码
func (rsp *Response) Code(code int) *Response {
	rsp.code = code
	return rsp
}

// Msg 设置响应消息
func (rsp *Response) Msg(msg string) *Response {
	rsp.message = msg
	return rsp
}

// Data 设置响应数据
func (rsp *Response) Data(data interface{}) *Response {
	rsp.data = data
	return rsp
}

// Raw 链式设置
func (rsp *Response) Raw(mix interface{}) *Response {
	switch i := mix.(type) {
	case int:
		rsp.Code(i)
	case error:
		rsp.Msg(i.Error())
	case string:
		rsp.Msg(i)
	default:
		rsp.Data(i)
	}

	return rsp
}

// RetJSON 返回json消息
func (rsp *Response) RetJSON() {
	data := map[string]interface{}{
		"code":    rsp.code,
		"message": rsp.message,
		"data":    rsp.data,
	}
	rsp.jsoner.Set("response_data", data)
	rsp.jsoner.JSON(http.StatusOK, data)
}

// NewResponse 生成一个响应结构体指针
func NewResponse(j JSONer) *Response {
	return &Response{jsoner: j}
}

// RetSucc 成功响应
func RetSucc(j JSONer, data interface{}) {
	NewResponse(j).Code(0).Msg("操作成功").Data(data).RetJSON()
}

// RetMixSucc 设置多个返回参数
func RetMixSucc(j JSONer, args ...interface{}) {
	rsp := NewResponse(j).Code(0).Msg("操作成功")
	for _, d := range args {
		rsp.Raw(d)
	}
	rsp.RetJSON()
}

// RetFail 失败响应
func RetFail(j JSONer, data interface{}) {
	NewResponse(j).Code(1).Msg("操作失败").Data(data).RetJSON()
}

// RetMixFail 设置多个返回参数
func RetMixFail(j JSONer, args ...interface{}) {
	rsp := NewResponse(j).Code(1).Msg("操作失败")
	for _, d := range args {
		rsp.Raw(d)
	}
	rsp.RetJSON()
}
