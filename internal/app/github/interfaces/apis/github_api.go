package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v41/github"
	"go.uber.org/zap"
	"io/ioutil"
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
	group := dc.ctx.GetRoute().Group("github")
	group.GET("/commits", wrapper(dc.GetRepoAllCommits))
	group.POST("/hook", wrapper(dc.GithubHook))
}

func (dc *GithubAPI) GithubHook(c *gin.Context) (interface{}, error) {
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	event, err := github.ParseWebHook(github.WebHookType(c.Request), payload)
	if err != nil {
		dc.logger.Error("could not parse webhook: err=%s\n", zap.Error(err))
		return nil, err
	}
	dc.logger.Debug("github hook event", zap.String("event", github.WebHookType(c.Request)))
	switch event.(type) {
	case *github.PushEvent:
	case *github.PullRequestEvent:
	case *github.WatchEvent:
	default:
		dc.logger.Warn("unknown event type", zap.String("event", github.WebHookType(c.Request)))
	}
	return "OK", nil
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
