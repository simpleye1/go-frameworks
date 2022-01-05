package context

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	github2 "github.com/google/go-github/v41/github"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"test/internal/pkg/cachestore"
	"test/internal/pkg/config"
	"test/internal/pkg/database"
	"test/internal/pkg/github"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/internal/pkg/redis"
	"test/internal/pkg/telemetry"
	"test/internal/pkg/transports/http"
)

func NewContext() context.Context {
	return context.Background()
}

// AppInfraContext Init需要放在Context里面承载，不然会被忽略
type AppInfraContext struct {
	MigrateInit   *migrate.Init
	TelemetryInit *telemetry.Init
	Config        *viper.Viper
	Log           *zap.Logger
	Route         *gin.Engine
	GormDB        *gorm.DB
	DB            *sql.DB
	CacheStore    cachestore.Store
	Context       context.Context
	GithubClient  *github2.Client
}

type InfraContext interface {
	GetConfig() *viper.Viper
	GetLog() *zap.Logger
	GetRoute() *gin.Engine
	GetGormDB() *gorm.DB
	GetDB() *sql.DB
	GetCacheStore() cachestore.Store
	GetContext() context.Context
	GetGithubClient() *github2.Client
}

func (a *AppInfraContext) GetConfig() *viper.Viper {
	return a.Config
}
func (a *AppInfraContext) GetLog() *zap.Logger {
	return a.Log
}
func (a *AppInfraContext) GetRoute() *gin.Engine {
	return a.Route
}
func (a *AppInfraContext) GetGormDB() *gorm.DB {
	return a.GormDB
}
func (a *AppInfraContext) GetDB() *sql.DB {
	return a.DB
}
func (a *AppInfraContext) GetCacheStore() cachestore.Store {
	return a.CacheStore
}
func (a *AppInfraContext) GetContext() context.Context {
	return a.Context
}
func (a *AppInfraContext) GetGithubClient() *github2.Client {
	return a.GithubClient
}

var ProviderSet = wire.NewSet(
	NewContext,
	wire.Struct(new(AppInfraContext), "*"),
	wire.Bind(new(InfraContext), new(*AppInfraContext)),
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	redis.ProviderSet,
	migrate.ProviderSet,
	http.ProviderSet,
	cachestore.ProviderSetRedis,
	telemetry.ProviderSet,
	github.ProviderSet,
)
