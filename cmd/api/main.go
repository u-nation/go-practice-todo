package main

import (
	"fmt"
	"github.com/u-nation/go-practice-todo/config"
	"github.com/u-nation/go-practice-todo/pkg/infrastructure/db"
	"github.com/u-nation/go-practice-todo/pkg/logger"
	"golang.org/x/xerrors"
)

func main() {
	apiConfig, err := config.NewAPIConfig()
	if err != nil {
		panic(err)
	}

	if err := logger.NewLogger(apiConfig.AppName, true); err != nil {
		panic(err)
	}

	server, err := initialize(apiConfig)
	if err != nil {
		fmt.Println(xerrors.Errorf("failed to initialize: %w", err))
		panic(err)
	}

	if err := db.MigrateDB(apiConfig.DB); err != nil {
		fmt.Println(xerrors.Errorf("failed to migrate: %w", err))
		panic(err)
	}

	server.run()
}
