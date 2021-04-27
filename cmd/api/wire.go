//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/u-nation/go-practice-todo/config"
	"github.com/u-nation/go-practice-todo/pkg/application/repository"
	"github.com/u-nation/go-practice-todo/pkg/infrastructure/db"
	rImpl "github.com/u-nation/go-practice-todo/pkg/infrastructure/repository"
	"github.com/u-nation/go-practice-todo/pkg/presentation"
)

func initialize(config *config.APIConfig) (*API, error) {
	wire.Build(ProvideApp)

	return nil, nil
}

var ProvideApp = wire.NewSet(
	wire.FieldsOf(new(*config.APIConfig), "DB"),
	ProvideRepository,
	//ProvideService,
	ProvidePresenter,
	//ProvideAPI,
	NewAPI,
)

var ProvideRepository = wire.NewSet(
	db.ProvideDB,
	wire.Struct(new(rImpl.TransactionRepositoryImpl), "*"),
	wire.Bind(new(repository.TransactionRepository), new(*rImpl.TransactionRepositoryImpl)),
	wire.Struct(new(rImpl.TodoRepositoryImpl), "*"),
	wire.Bind(new(repository.TodoRepository), new(*rImpl.TodoRepositoryImpl)),
	wire.Struct(new(rImpl.HealthCheckRepositoryImpl), "*"),
	wire.Bind(new(repository.HealthCheckRepository), new(*rImpl.HealthCheckRepositoryImpl)),
)

var ProvidePresenter = wire.NewSet(
	presentation.NewApplicationController,
	presentation.HealthCheckControllerSet,
)
