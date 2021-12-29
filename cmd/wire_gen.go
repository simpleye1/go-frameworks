// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	app2 "test/internal/app"
	"test/internal/app/module1/application"
	"test/internal/app/module1/domain/services"
	"test/internal/app/module1/infrastructure/repos"
	"test/internal/app/module1/interfaces"
	"test/internal/app/module1/interfaces/apis"
	"test/internal/pkg"
	"test/internal/pkg/app"
	"test/internal/pkg/config"
	"test/internal/pkg/database"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/internal/pkg/transports/http"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	appOptions, err := app2.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	migrationOptions, err := migrate.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	gormDB, err := migrate.Migrate(viper, databaseOptions, migrationOptions, db, logger)
	if err != nil {
		return nil, err
	}
	postgresDetailRepository := repos.NewPostgresDetailsRepository(logger, gormDB)
	postgresUserRepository := repos.NewPostgresUserRepository(logger, gormDB)
	userDetailServiceImpl := services.NewUserDetailServiceImpl(logger, postgresDetailRepository, postgresUserRepository)
	userDetailApplication := application.NewDetailsApplication(logger, userDetailServiceImpl)
	v := interfaces.NewAPIS(logger, userDetailApplication)
	initControllers := apis.CreateInitControllersFn(v...)
	engine := http.NewRouter(httpOptions, logger, initControllers)
	server, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	context := &app.Context{
		Config: viper,
		Log:    logger,
		Engine: engine,
		Server: server,
		GormDB: gormDB,
		DB:     db,
	}
	appApplication, err := app2.NewApp(appOptions, context, logger, server)
	if err != nil {
		return nil, err
	}
	return appApplication, nil
}

// wire.go:

var providerSet = wire.NewSet(pkg.ProviderSet, app2.ProviderSet)
