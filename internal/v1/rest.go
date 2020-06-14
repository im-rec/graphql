package v1

import (
	"github.com/labstack/echo"
	"github.com/rianekacahya/graphql/internal/v1/config"
	"github.com/rianekacahya/graphql/internal/v1/module/hello"
	"github.com/rianekacahya/graphql/internal/v1/module/hello/transport/rest"
	"github.com/rianekacahya/graphql/pkg/echoserver"
	"net/http"
)

func StartRest() {
	// start all infrastructure dependencies
	bootstrap()

	// run echo server
	echoserver.InitServer(config.GetConfig().AppDebug)

	// set healthcheck endpoint
	echoserver.GetServer().GET("/v1/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// init Service
	rest.NewRest(hello.Initialize())

	// start server
	go echoserver.StartServer(
		config.GetConfig().ServerRestPort,
		config.GetConfig().ServerRestWriteTimeout,
		config.GetConfig().ServerRestReadTimeout,
		config.GetConfig().ServerRestIdleTimeout,
	)

	// shutdown server gracefully
	echoserver.Shutdown()
}
