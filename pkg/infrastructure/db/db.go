package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/u-nation/go-practice-todo/config"
	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"

	// imported drive
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	// imported driver file
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ProvideDB(config *config.DBConfig) (*gorm.DB, error) {
	// Logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	writerDB, err := gorm.Open(mysql.Open(config.Writer.GetConnectionUrl()), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, xerrors.Errorf("failed to connect db: %w", err)
	}

	if err = writerDB.Use(
		dbresolver.Register(
			dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(config.Reader.GetConnectionUrl())},
				Policy:   dbresolver.RandomPolicy{},
			},
		).SetMaxIdleConns(config.MaxIdleConns).SetMaxOpenConns(config.MaxOpenConns),
	); err != nil {
		return nil, err
	}

	return writerDB, nil
}

// MigrateDB DBのマイグレーションを実行.
func MigrateDB(config *config.DBConfig) error {
	// Run the migrations
	m, err := migrate.New(
		fmt.Sprintf("file://%s", config.MigrationDir),
		fmt.Sprintf("mysql://%s", config.Writer.GetConnectionUrl()),
	)

	if err != nil {
		return xerrors.Errorf("failed create migrate: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return xerrors.Errorf("failed run migrate: %w", err)
	}

	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return xerrors.Errorf("migrate source error: %w", srcErr)
	}

	if dbErr != nil {
		return xerrors.Errorf("migrate database error: %w", dbErr)
	}

	return nil
}