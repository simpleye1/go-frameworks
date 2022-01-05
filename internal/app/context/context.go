package context

import (
	"github.com/google/wire"
	"test/internal/app/github/application"
	services2 "test/internal/app/github/application/services"
	"test/internal/app/github/domain/clients"
	repos2 "test/internal/app/github/domain/repos"
	"test/internal/app/github/domain/services"
	clients2 "test/internal/app/github/infrastructure/clients"
	"test/internal/app/github/infrastructure/repos"
	"test/internal/app/github/interfaces/apis"
	"test/internal/pkg/context"
)

type AppContext struct {
	context.InfraContext

	*apis.GithubAPI

	*application.GithubApplication

	repos2.UserRepository
	repos2.DetailRepository

	services2.GithubService
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppContext), "*"),
	// API
	APIProviderSet,
	// Application
	ApplicationProviderSet,
	// Client
	ClientProviderSet,
	// Service
	ServiceProviderSet,
	// Repo
	RepoProviderSet,
)

var APIProviderSet = wire.NewSet(
	apis.NewAPI,
	apis.NewGithubAPI,
)

var ApplicationProviderSet = wire.NewSet(
	application.NewUserDetailsApplication,
)

var ClientProviderSet = wire.NewSet(
	clients2.NewGithubClientImpl,
	wire.Bind(new(clients.GithubClient), new(*clients2.GithubClientImpl)),
)

var ServiceProviderSet = wire.NewSet(
	services.NewUserDetailServiceImpl,
	wire.Bind(new(services2.GithubService), new(*services.GithubServiceImpl)),
)

var RepoProviderSet = wire.NewSet(
	repos.NewPostgresDetailsRepository,
	repos.NewPostgresUserRepository,
	wire.Bind(new(repos2.UserRepository), new(*repos.PostgresUserRepository)),
	wire.Bind(new(repos2.DetailRepository), new(*repos.PostgresDetailRepository)),
)
