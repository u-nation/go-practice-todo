package presentation

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/u-nation/go-practice-todo/config"
	"github.com/u-nation/go-practice-todo/pkg/util"
)

type ApplicationController struct {
	config *config.APIConfig
}

func NewApplicationController(config *config.APIConfig) ApplicationController {
	return ApplicationController{
		config,
	}
}

func (ac ApplicationController) CreateContext(c *echo.Context) context.Context {
	ctx := (*c).Request().Context()

	if requestID, exists := (*c).Get(util.RequestIDKey).(string); exists {
		ctx = util.WithRequestID(ctx, requestID)
	}

	return ctx
}
