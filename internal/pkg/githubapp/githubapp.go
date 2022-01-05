package githubapp

import (
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v39/github"
	"github.com/google/wire"
	"log"
	"net/http"
)

func NewGithubApp() (*github.Client, error) {
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, 162626, 21954910, "internal/pkg/githubapp/xingbo-go.2022-01-04.private-key.pem")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	client := github.NewClient(&http.Client{Transport: itr})
	return client, nil
}

var ProviderSet = wire.NewSet(NewGithubApp)
