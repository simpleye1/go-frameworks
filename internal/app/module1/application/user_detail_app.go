package application

import (
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
	"go.uber.org/zap"
	"test/internal/app/module1/domain/services"
)

type UserDetailApplication struct {
	logger  *zap.Logger
	service services.UserDetailService
	client  *github.Client
}

func NewUserDetailsApplication(logger *zap.Logger, s services.UserDetailService, client *github.Client) *UserDetailApplication {
	u := &UserDetailApplication{
		logger:  logger,
		service: s,
		client:  client,
	}
	return u
}

func (da *UserDetailApplication) GetUserDetail(id uint64) (*services.UserDetail, error) {
	return da.service.GetUserDetail(id)
}

func (da *UserDetailApplication) GetCommits(user string, repo string) []*github.RepositoryCommit {
	commits, _, err := da.client.Repositories.ListCommits(context.Background(), user, repo, nil)
	if err != nil {
		fmt.Errorf("not find")
	}
	return commits
}
