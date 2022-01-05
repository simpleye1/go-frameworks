package clients

type GithubClient interface {
	GetRepoAllCommits(owner string, repo string, sha string, page int) ([]*RepoCommit, error)
}
type RepoCommit struct {
	SHA       string
	Commit    string
	Committer string
	HTMLURL   string
}
