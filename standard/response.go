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

// RetFail 失败响应
func RetFail(j JSONer, data interface{}) {
	NewResponse(j).Code(1).Msg("操作失败").Data(data).RetJSON()
}
