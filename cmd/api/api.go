package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/u-nation/go-practice-todo/config"
	"github.com/u-nation/go-practice-todo/pkg/presentation"
	"github.com/u-nation/go-practice-todo/pkg/util"
	"golang.org/x/net/netutil"
	"log"
	"net"
	"net/http"
	"os"
)

type API struct {
	Config *config.APIConfig
	Server *echo.Echo
}

func NewAPI(
	config *config.APIConfig,
	healthcheck presentation.HealthCheckController,
) *API {
	api := &API{
		Config: config,
		Server: echo.New(),
	}
	api.setListener()
	api.setMiddleware()

	api.Server.GET("/healthcheck", healthcheck.HealthCheck)
	api.Server.GET("/version", func(ectx echo.Context) error { return ectx.JSON(http.StatusOK, os.Getenv("API_VERSION")) })

	//v1 := api.Server.Group("/v1")

	return api
}

func (a *API) setListener() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	a.Server.Listener = netutil.LimitListener(ln, a.Config.MaxConns)
}

func (a *API) setMiddleware() {

	a.Server.Use(middleware.Recover())
	a.Server.Use(middleware.CORS())
	a.Server.Use(util.RequestID())
	a.Server.Use(middleware.LoggerWithConfig(util.LoggerConfig()))
}

func (a *API) run() {
	a.Server.Logger.Fatal(a.Server.Start(":5000"))
}
