package application

import (
	"go.uber.org/zap"
	"test/internal/app/github/application/services"
	"test/internal/app/github/domain/clients"
)

type GithubApplication struct {
	logger  *zap.Logger
	service services.GithubService
}

func NewUserDetailsApplication(logger *zap.Logger, s services.GithubService) *GithubApplication {
	u := &GithubApplication{
		logger:  logger,
		service: s,
	}
	return u
}

func (da *GithubApplication) GetRepoAllCommits(owner string, repo string, sha string, page int) ([]*clients.RepoCommit, error) {
	return da.service.GetRepoAllCommits(owner, repo, sha, page)
}
