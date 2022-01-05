package exceptions

type AppException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

// 实现接口
func (e *AppException) Error() string {
	return e.Msg
}

func NewAppException(code int, errorCode int, msg string) *AppException {
	return &AppException{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}
