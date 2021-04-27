package presentation

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/u-nation/go-practice-todo/pkg/application/repository"
	"golang.org/x/xerrors"
	"net/http"
)

type HealthCheckController struct {
	ApplicationController
	repository.HealthCheckRepository
}

var HealthCheckControllerSet = wire.NewSet(
	wire.Struct(new(HealthCheckController), "*"),
)

func (hc *HealthCheckController) HealthCheck(ectx echo.Context) error {
	ctx := hc.CreateContext(&ectx)

	if err := hc.HealthCheckRepository.Ping(ctx); err != nil {
		return xerrors.Errorf("health check error: %w", err)
	}

	return ectx.JSON(http.StatusOK, "OK")
}
