package apis

import (
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
	"test/internal/app/module1/application"
	"test/internal/app/module1/interfaces/exceptions"
	"time"
)

type UserDetailAPI struct {
	API
	application *application.UserDetailApplication
}

func NewUserDetailAPI(api *API, a *application.UserDetailApplication) *UserDetailAPI {
	v := &UserDetailAPI{
		API:         *api,
		application: a,
	}
	v.Init()
	return v
}

func (dc *UserDetailAPI) Init() {
	dc.ctx.GetRoute().GET("/detail", cache.CacheByRequestURI(dc.ctx.GetCacheStore(), 2*time.Second), wrapper(dc.GetUserDetail))
	dc.ctx.GetRoute().GET("/commits/:user/:repo", cache.CacheByRequestURI(dc.ctx.GetCacheStore(), 2*time.Second), wrapper(dc.GetCommits))
}

func (dc *UserDetailAPI) GetUserDetail(c *gin.Context) (interface{}, error) {
	param := struct {
		ID uint64 `form:"id" binding:"required"`
	}{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		return nil, exceptions.ParameterError(err.Error())
	}
	p, err := dc.application.GetUserDetail(param.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (dc *UserDetailAPI) GetCommits(c *gin.Context) (interface{}, error) {
	user, _ := c.Params.Get("user")
	repo, _ := c.Params.Get("repo")
	//todo 校验，错误验证
	commits := dc.application.GetCommits(user, repo)
	return commits, nil
}
