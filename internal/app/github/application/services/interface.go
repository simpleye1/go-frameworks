package services

import (
	"test/internal/app/github/domain/clients"
)

type GithubService interface {
	GetRepoAllCommits(owner string, repo string, sha string, page int) ([]*clients.RepoCommit, error)
}
