package util

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerConfig() middleware.LoggerConfig {
	ignorePath := []string{"/healthcheck", "/version", "/demo/error"}

	return middleware.LoggerConfig{
		Skipper: func(ctx echo.Context) bool {
			for _, path := range ignorePath {
				if ctx.Path() == path {
					return true
				}
			}
			return false
		},
	}
}

