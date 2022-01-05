package clients

import (
	"context"
	"github.com/google/go-github/v41/github"
	"test/internal/app/github/domain/clients"
	"test/internal/app/github/domain/exceptions"
)

type GithubClientImpl struct {
	client *github.Client
	ctx    context.Context
}

func NewGithubClientImpl(client *github.Client, ctx context.Context) *GithubClientImpl {
	return &GithubClientImpl{
		client: client,
		ctx:    ctx,
	}
}

func (g GithubClientImpl) GetRepoAllCommits(owner string, repo string, sha string, page int) ([]*clients.RepoCommit, error) {
	commits, r, err := g.client.Repositories.ListCommits(g.ctx, owner, repo, &github.CommitsListOptions{
		SHA:    sha,
		Path:   "",
		Author: "",
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: 30,
		},
	})
	if r.StatusCode != 200 {
		return nil, exceptions.BusinessError("github api get error")
	}
	if err != nil {
		return nil, err
	}
	var repoCommits []*clients.RepoCommit
	for _, commit := range commits {
		rc := &clients.RepoCommit{
			SHA:     *commit.SHA,
			HTMLURL: *commit.HTMLURL,
		}
		if commit.Commit != nil {
			rc.Commit = *commit.Commit.Message
			rc.Committer = *commit.Commit.Author.Name
		}
		repoCommits = append(repoCommits, rc)
	}
	return repoCommits, nil
}
