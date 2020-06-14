package v1

import (
	"github.com/labstack/echo"
	"github.com/rianekacahya/graphql/internal/v1/config"
	"github.com/rianekacahya/graphql/pkg/echoserver"
	"net/http"
)

func StartGraphQL() {
	// start all infrastructure dependencies
	bootstrap()

	// run echo server
	echoserver.InitServer(config.GetConfig().AppDebug)

	// set healthcheck endpoint
	echoserver.GetServer().GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// start server
	go echoserver.StartServer(
		config.GetConfig().ServerGraphqlPort,
		config.GetConfig().ServerGraphqlWriteTimeout,
		config.GetConfig().ServerGraphqlReadTimeout,
		config.GetConfig().ServerGraphqlIdleTimeout,
	)

	// shutdown server gracefully
	echoserver.Shutdown()
}
