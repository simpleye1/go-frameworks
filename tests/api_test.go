package tests

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test/internal/app/github/domain/clients"
	"test/mocks"
	"testing"
)

func TestGithubApi(t *testing.T) {
	flag.Parse()
	us := new(mocks.GithubService)
	var rcs = []*clients.RepoCommit{
		&clients.RepoCommit{
			SHA:       "sha",
			Commit:    "commit",
			Committer: "user",
			HTMLURL:   "html",
		},
	}

	us.On("GetRepoAllCommits",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
	).Return(rcs, nil)
	api, err := CreateGithubAPI(*resourcesPath, us)
	if err != nil {
		t.Fatalf("get userDetail api  error,%+v", err)
	}
	resp := callAPI(api.API, "GET", "/commits?owner=owner&repo=repo&page=0", nil)
	assert.Equal(t, 200, resp.StatusCode)
}
