package github

import (
	"github.com/google/go-github/v41/github"
	"github.com/google/wire"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"net/http"
)

func NewGithub() (*github.Client, error) {
	client := github.NewClient(newMockGithubClient())
	return client, nil
}

func newMockGithubClient() *http.Client {
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetReposCommitsByOwnerByRepo,
			[]*github.RepositoryCommit{
				{
					SHA: github.String("SHA"),
					Commit: &github.Commit{
						Author: &github.CommitAuthor{
							Name: github.String("Name"),
						},
						Message: github.String("Message"),
					},
					Author:    nil,
					Committer: nil,
					HTMLURL:   github.String("HTMLURL"),
				},
			},
		),
	)
	return mockedHTTPClient
}

var ProviderSet = wire.NewSet(NewGithub)
