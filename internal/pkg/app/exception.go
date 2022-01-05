package app

import "net/http"

type Exception struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

// 实现接口
func (e *Exception) Error() string {
	return e.Msg
}

func NewAppException(code int, errorCode int, msg string) *Exception {
	return &Exception{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}

const (
	SERVER_ERROR    = 1000 // 系统错误
	NOT_FOUND       = 1001 // 401错误
	UNKNOWN_ERROR   = 1002 // 未知错误
	PARAMETER_ERROR = 1003 // 参数错误
	AUTH_ERROR      = 1004 // 错误
	BUSINESS_ERROR  = 2000 // 业务错误
)

// ServerError 500 错误处理
func ServerError() *Exception {
	return NewAppException(http.StatusInternalServerError, SERVER_ERROR, http.StatusText(http.StatusInternalServerError))
}

// NotFound 404 错误
func NotFound() *Exception {
	return NewAppException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// UnknownError 未知错误
func UnknownError(message string) *Exception {
	return NewAppException(http.StatusForbidden, UNKNOWN_ERROR, message)
}

// ParameterError 参数错误
func ParameterError(message string) *Exception {
	return NewAppException(http.StatusBadRequest, PARAMETER_ERROR, message)
}

// BusinessError 业务错误
func BusinessError(message string) *Exception {
	return NewAppException(http.StatusBadRequest, BUSINESS_ERROR, message)
}
