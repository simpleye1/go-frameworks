package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"test/internal/pkg/app"
	"test/internal/pkg/context"
)

type API struct {
	logger *zap.Logger
	ctx    context.InfraContext
}

func (a *API) GetInfraContext() context.InfraContext {
	return a.ctx
}

func NewAPI(logger *zap.Logger, ctx context.InfraContext) *API {
	return &API{
		logger: logger,
		ctx:    ctx,
	}
}

type HandlerFunc func(c *gin.Context) (interface{}, error)

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err  error
			body interface{}
		)
		body, err = handler(c)
		if err != nil {
			var apiException *app.Exception
			if h, ok := err.(*app.Exception); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = app.UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = app.UnknownError(e.Error())
				}
			} else {
				apiException = app.ServerError()
			}
			apiException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(apiException.Code, apiException)
			return
		} else {
			c.JSON(http.StatusOK, body)
		}
	}
}
