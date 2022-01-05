package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithubService_GetRepoAllCommits(t *testing.T) {
	background := setUp()
	rcs, err := background.GithubService.GetRepoAllCommits("foobar", "foobar", "", 0)
	if err != nil {
		t.Fatalf("userDetail service get userDetail error,%+v", err)
	}
	assert.True(t, len(rcs) > 0)
	assert.NotNil(t, rcs[0].Commit)
	assert.NotNil(t, rcs[0].SHA)
	assert.NotNil(t, rcs[0].HTMLURL)
	assert.NotNil(t, rcs[0].Committer)
}
