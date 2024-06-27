package res

import (
	"github.com/gin-gonic/gin"
	"gvb_server/uploads/utils"
	"net/http"
)

const (
	Success = 0
	Error   = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

// c *gin.Context参数是一个指向gin.Context结构体实例的指针。
// 这个c代表当前HTTP请求的上下文，它包含了关于这个请求的所有信息，以及用于构造响应的各种方法。
func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "Success", c)
}

func OKWithList(list any, count int64, c *gin.Context) {
	OKWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func OKWith(c *gin.Context) {
	Result(Success, map[string]any{}, "success", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

// map: 实际上是告诉Result函数我们这次响应不需要返回具体的数据，只需要返回一个消息（msg参数指定的内容）即可。
// 这个空的映射表明响应体 （response body）中不包含数据字段，或者数据字段为空
// code 是状态码 需要去查一下他的msg
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[ErrorCode(code)]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "unknown error", c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
