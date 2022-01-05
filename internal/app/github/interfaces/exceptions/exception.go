package exceptions

import (
	"net/http"
	"test/internal/app/github/infrastructure/exceptions"
)

const (
	SERVER_ERROR    = 1000 // 系统错误
	NOT_FOUND       = 1001 // 401错误
	UNKNOWN_ERROR   = 1002 // 未知错误
	PARAMETER_ERROR = 1003 // 参数错误
	AUTH_ERROR      = 1004 // 错误
)

// ServerError 500 错误处理
func ServerError() *exceptions.AppException {
	return exceptions.NewAppException(http.StatusInternalServerError, SERVER_ERROR, http.StatusText(http.StatusInternalServerError))
}

// NotFound 404 错误
func NotFound() *exceptions.AppException {
	return exceptions.NewAppException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// UnknownError 未知错误
func UnknownError(message string) *exceptions.AppException {
	return exceptions.NewAppException(http.StatusForbidden, UNKNOWN_ERROR, message)
}

// ParameterError 参数错误
func ParameterError(message string) *exceptions.AppException {
	return exceptions.NewAppException(http.StatusBadRequest, PARAMETER_ERROR, message)
}
