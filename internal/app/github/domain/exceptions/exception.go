package exceptions

import (
	"net/http"
	"test/internal/app/github/infrastructure/exceptions"
)

const (
	BUSINESS_ERROR = 2000 // 业务错误
)

// BusinessError 业务错误
func BusinessError(message string) *exceptions.AppException {
	return exceptions.NewAppException(http.StatusBadRequest, BUSINESS_ERROR, message)
}
