package cuckoo_error

import (
	"net/http"
)

type IErrorStatus interface {
	Code() int
	StatusCode() int
	Message() string
}

type _ErrorStatus struct {
	code    int
	message string
}

func (err *_ErrorStatus) Code() int {
	return err.code
}

func (err *_ErrorStatus) Message() string {
	return err.message
}

func (err *_ErrorStatus) StatusCode() int {
	switch err.Code() {
	case Success.Code():
		return http.StatusOK
	case InternalError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UserNotExist.Code(), UserAuthFailed.Code(), UserAlreadyExist.Code():
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

func NewErrorStatus(code int, msg string) *_ErrorStatus {
	return &_ErrorStatus{
		code:    code,
		message: msg,
	}
}

var (
	Success              = NewErrorStatus(0, "成功")
	InternalError        = NewErrorStatus(10000, "服务内部错误")
	InvalidParams        = NewErrorStatus(10001, "客户端请求参数错误")
	UserNotExist         = NewErrorStatus(10002, "账号不存在")
	UserAlreadyExist     = NewErrorStatus(10003, "账号已存在")
	UserAuthFailed       = NewErrorStatus(10004, "用户名或密码错误")
	UserAuthTokenError   = NewErrorStatus(10005, "token错误")
	UserAuthTokenExpired = NewErrorStatus(10006, "token失效")
)
