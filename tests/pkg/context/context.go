package context

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"test/internal/pkg/cachestore"
	"test/internal/pkg/config"
	context2 "test/internal/pkg/context"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/tests/pkg/database"
	"test/tests/pkg/redis"
	"test/tests/pkg/transports/http"
)

func NewContext() context.Context {
	return context.Background()
}

// TestInfraContext Init需要放在Context里面承载，不然会被忽略
type TestInfraContext struct {
	MigrateInit *migrate.Init
	Config      *viper.Viper
	Log         *zap.Logger
	Route       *gin.Engine
	GormDB      *gorm.DB
	DB          *sql.DB
	CacheStore  cachestore.Store
	Context     context.Context
	GithubApp   *github.Client
}

func (a *TestInfraContext) GetGithubApp() *github.Client {
	return a.GithubApp
}

func (a *TestInfraContext) GetConfig() *viper.Viper {
	return a.Config
}
func (a *TestInfraContext) GetLog() *zap.Logger {
	return a.Log
}
func (a *TestInfraContext) GetRoute() *gin.Engine {
	return a.Route
}
func (a *TestInfraContext) GetGormDB() *gorm.DB {
	return a.GormDB
}
func (a *TestInfraContext) GetDB() *sql.DB {
	return a.DB
}
func (a *TestInfraContext) GetCacheStore() cachestore.Store {
	return a.CacheStore
}
func (a *TestInfraContext) GetContext() context.Context {
	return a.Context
}

var ProviderSet = wire.NewSet(
	NewContext,
	wire.Struct(new(TestInfraContext), "*"),
	wire.Bind(new(context2.InfraContext), new(*TestInfraContext)),
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	redis.ProviderSet,
	migrate.ProviderSet,
	http.ProviderSet,
	cachestore.ProviderSetRedis,
)

type TestMockAPIInfraContext struct {
	Config     *viper.Viper
	Log        *zap.Logger
	Route      *gin.Engine
	CacheStore cachestore.Store
}

func (a *TestMockAPIInfraContext) GetConfig() *viper.Viper {
	return a.Config
}
func (a *TestMockAPIInfraContext) GetLog() *zap.Logger {
	return a.Log
}
func (a *TestMockAPIInfraContext) GetRoute() *gin.Engine {
	return a.Route
}
func (a *TestMockAPIInfraContext) GetGormDB() *gorm.DB {
	return nil
}
func (a *TestMockAPIInfraContext) GetDB() *sql.DB {
	return nil
}
func (a *TestMockAPIInfraContext) GetCacheStore() cachestore.Store {
	return a.CacheStore
}
func (a *TestMockAPIInfraContext) GetContext() context.Context {
	return nil
}

var APIMockProviderSet = wire.NewSet(
	NewContext,
	wire.Struct(new(TestMockAPIInfraContext), "*"),
	wire.Bind(new(context2.InfraContext), new(*TestMockAPIInfraContext)),
	log.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	cachestore.ProviderSetMemory,
)
