//go:generate go run github.com/google/wire/cmd/wire
// +build wireinject

package tests

import (
	"github.com/google/wire"
	"test/internal/app"
	"test/internal/app/context"
	"test/internal/app/github/application/services"
	"test/internal/app/github/interfaces/apis"
	"test/tests/pkg"
	"test/tests/pkg/testcontainer"
)

var ProviderSet = wire.NewSet(
	app.ProviderSet,
	pkg.ProviderSet,
)

func CreateBackground(cf string) (*testcontainer.Background, func(), error) {
	panic(wire.Build(ProviderSet))
}

func CreateGithubAPI(cf string, s services.GithubService) (*apis.GithubAPI, error) {
	panic(wire.Build(pkg.APIMockProviderSet, context.APIProviderSet, context.ApplicationProviderSet))
}
