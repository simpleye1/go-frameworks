package context

import (
	"github.com/google/wire"
	"test/internal/app/github/application"
	servicesDef "test/internal/app/github/application/services"
	clientsDef "test/internal/app/github/domain/clients"
	reposDef "test/internal/app/github/domain/repos"
	"test/internal/app/github/domain/services"
	"test/internal/app/github/infrastructure/clients"
	"test/internal/app/github/infrastructure/repos"
	"test/internal/app/github/interfaces/apis"
	"test/internal/pkg/context"
)

type AppContext struct {
	context.InfraContext

	*apis.GithubAPI

	*application.GithubApplication

	reposDef.UserRepository
	reposDef.DetailRepository

	servicesDef.GithubService
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
	clients.NewGithubClientImpl,
	wire.Bind(new(clientsDef.GithubClient), new(*clients.GithubClientImpl)),
)

var ServiceProviderSet = wire.NewSet(
	services.NewUserDetailServiceImpl,
	wire.Bind(new(servicesDef.GithubService), new(*services.GithubServiceImpl)),
)

var RepoProviderSet = wire.NewSet(
	repos.NewPostgresDetailsRepository,
	repos.NewPostgresUserRepository,
	wire.Bind(new(reposDef.UserRepository), new(*repos.PostgresUserRepository)),
	wire.Bind(new(reposDef.DetailRepository), new(*repos.PostgresDetailRepository)),
)
