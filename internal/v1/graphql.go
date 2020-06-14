package v1

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/rianekacahya/graphql/internal/v1/config"
	"github.com/rianekacahya/graphql/internal/v1/module/hello"
	"github.com/rianekacahya/graphql/internal/v1/module/hello/transport/graphql"
	"github.com/rianekacahya/graphql/pkg/echoserver"
	"net/http"
)

func StartGraphQL() {
	// start all infrastructure dependencies
	bootstrap()

	// run echo server
	echoserver.InitServer(config.GetConfig().AppDebug)

	// set graphQL playground
	echoserver.GetServer().GET("/", echo.WrapHandler(playground.Handler("GraphQL Playground", "/")))

	// set healthcheck endpoint
	echoserver.GetServer().GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// init Service
	graphql.NewGraphQL(hello.Initialize())

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
