package services

import (
	"go.uber.org/zap"
	"test/internal/app/github/domain/clients"
)

type GithubServiceImpl struct {
	logger       *zap.Logger
	githubClient clients.GithubClient
}

func NewUserDetailServiceImpl(
	logger *zap.Logger,
	githubClient clients.GithubClient,
) *GithubServiceImpl {
	u := &GithubServiceImpl{
		logger:       logger.With(zap.String("type", "GithubServiceImpl")),
		githubClient: githubClient,
	}
	return u
}

func (g GithubServiceImpl) GetRepoAllCommits(owner string, repo string, sha string, page int) ([]*clients.RepoCommit, error) {
	return g.githubClient.GetRepoAllCommits(owner, repo, sha, page)
}
