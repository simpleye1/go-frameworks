package apis

import (
	"github.com/gin-gonic/gin"
	"test/internal/app/github/application"
	"test/internal/app/github/interfaces/exceptions"
)

type GithubAPI struct {
	API
	application *application.GithubApplication
}

func NewGithubAPI(api *API, a *application.GithubApplication) *GithubAPI {
	v := &GithubAPI{
		API:         *api,
		application: a,
	}
	v.Init()
	return v
}

func (dc *GithubAPI) Init() {
	dc.ctx.GetRoute().GET("/commits", wrapper(dc.GetRepoAllCommits))
}

func (dc *GithubAPI) GetRepoAllCommits(c *gin.Context) (interface{}, error) {
	param := struct {
		Owner string `form:"owner" binding:"required"`
		Repo  string `form:"repo" binding:"required"`
		SHA   string `form:"sha"`
		Page  int    `form:"page"`
	}{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		return nil, exceptions.ParameterError(err.Error())
	}
	p, err := dc.application.GetRepoAllCommits(param.Owner, param.Repo, param.SHA, param.Page)
	if err != nil {
		return nil, err
	}
	return p, nil
}
